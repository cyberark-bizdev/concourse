package builder

import (
	"io"
	"io/ioutil"
	"strings"
	"time"
	"unicode/utf8"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/event"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/vars"
)

func NewDelegateFactory() *delegateFactory {
	return &delegateFactory{}
}

type delegateFactory struct{}

func (delegate *delegateFactory) GetDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker) exec.GetDelegate {
	return NewGetDelegate(build, planID, credVarsTracker, clock.NewClock())
}

func (delegate *delegateFactory) PutDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker) exec.PutDelegate {
	return NewPutDelegate(build, planID, credVarsTracker, clock.NewClock())
}

func (delegate *delegateFactory) TaskDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker) exec.TaskDelegate {
	return NewTaskDelegate(build, planID, credVarsTracker, clock.NewClock())
}

func (delegate *delegateFactory) CheckDelegate(check db.Check, planID atc.PlanID, credVarsTracker vars.CredVarsTracker) exec.CheckDelegate {
	return NewCheckDelegate(check, planID, credVarsTracker, clock.NewClock())
}

func (delegate *delegateFactory) BuildStepDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker) exec.BuildStepDelegate {
	return NewBuildStepDelegate(build, planID, credVarsTracker, clock.NewClock())
}

func NewGetDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker, clock clock.Clock) exec.GetDelegate {
	return &getDelegate{
		BuildStepDelegate: NewBuildStepDelegate(build, planID, credVarsTracker, clock),

		eventOrigin: event.Origin{ID: event.OriginID(planID)},
		build:       build,
		clock:       clock,
	}
}

type getDelegate struct {
	exec.BuildStepDelegate

	build       db.Build
	eventOrigin event.Origin
	clock       clock.Clock
}

func (d *getDelegate) Initializing(logger lager.Logger) {
	err := d.build.SaveEvent(event.InitializeGet{
		Origin: d.eventOrigin,
		Time:   time.Now().Unix(),
	})
	if err != nil {
		logger.Error("failed-to-save-initialize-get-event", err)
		return
	}

	logger.Info("initializing")
}

func (d *getDelegate) Starting(logger lager.Logger) {
	err := d.build.SaveEvent(event.StartGet{
		Time:   time.Now().Unix(),
		Origin: d.eventOrigin,
	})
	if err != nil {
		logger.Error("failed-to-save-start-get-event", err)
		return
	}

	logger.Info("starting")
}

func (d *getDelegate) Finished(logger lager.Logger, exitStatus exec.ExitStatus, info exec.VersionInfo) {
	// PR#4398: close to flush stdout and stderr
	d.Stdout().(io.Closer).Close()
	d.Stderr().(io.Closer).Close()

	err := d.build.SaveEvent(event.FinishGet{
		Origin:          d.eventOrigin,
		Time:            d.clock.Now().Unix(),
		ExitStatus:      int(exitStatus),
		FetchedVersion:  info.Version,
		FetchedMetadata: info.Metadata,
	})
	if err != nil {
		logger.Error("failed-to-save-finish-get-event", err)
		return
	}

	logger.Info("finished", lager.Data{"exit-status": exitStatus})
}

func (d *getDelegate) UpdateVersion(log lager.Logger, plan atc.GetPlan, info exec.VersionInfo) {
	logger := log.WithData(lager.Data{
		"pipeline-name": d.build.PipelineName(),
		"pipeline-id":   d.build.PipelineID()},
	)

	pipeline, found, err := d.build.Pipeline()
	if err != nil {
		logger.Error("failed-to-find-pipeline", err)
		return
	}

	if !found {
		logger.Debug("pipeline-not-found")
		return
	}

	resource, found, err := pipeline.Resource(plan.Resource)
	if err != nil {
		logger.Error("failed-to-find-resource", err)
		return
	}

	if !found {
		logger.Debug("resource-not-found")
		return
	}

	_, err = resource.UpdateMetadata(
		info.Version,
		db.NewResourceConfigMetadataFields(info.Metadata),
	)
	if err != nil {
		logger.Error("failed-to-save-resource-config-version-metadata", err)
		return
	}
}

func NewPutDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker, clock clock.Clock) exec.PutDelegate {
	return &putDelegate{
		BuildStepDelegate: NewBuildStepDelegate(build, planID, credVarsTracker, clock),

		eventOrigin: event.Origin{ID: event.OriginID(planID)},
		build:       build,
		clock:       clock,
	}
}

type putDelegate struct {
	exec.BuildStepDelegate

	build       db.Build
	eventOrigin event.Origin
	clock       clock.Clock
}

func (d *putDelegate) Initializing(logger lager.Logger) {
	err := d.build.SaveEvent(event.InitializePut{
		Origin: d.eventOrigin,
		Time:   time.Now().Unix(),
	})
	if err != nil {
		logger.Error("failed-to-save-initialize-put-event", err)
		return
	}

	logger.Info("initializing")
}

func (d *putDelegate) Starting(logger lager.Logger) {
	err := d.build.SaveEvent(event.StartPut{
		Time:   time.Now().Unix(),
		Origin: d.eventOrigin,
	})
	if err != nil {
		logger.Error("failed-to-save-start-put-event", err)
		return
	}

	logger.Info("starting")
}

func (d *putDelegate) Finished(logger lager.Logger, exitStatus exec.ExitStatus, info exec.VersionInfo) {
	// PR#4398: close to flush stdout and stderr
	d.Stdout().(io.Closer).Close()
	d.Stderr().(io.Closer).Close()

	err := d.build.SaveEvent(event.FinishPut{
		Origin:          d.eventOrigin,
		Time:            d.clock.Now().Unix(),
		ExitStatus:      int(exitStatus),
		CreatedVersion:  info.Version,
		CreatedMetadata: info.Metadata,
	})
	if err != nil {
		logger.Error("failed-to-save-finish-put-event", err)
		return
	}

	logger.Info("finished", lager.Data{"exit-status": exitStatus, "version-info": info})
}

func (d *putDelegate) SaveOutput(log lager.Logger, plan atc.PutPlan, source atc.Source, resourceTypes atc.VersionedResourceTypes, info exec.VersionInfo) {
	logger := log.WithData(lager.Data{
		"step":          plan.Name,
		"resource":      plan.Resource,
		"resource-type": plan.Type,
		"version":       info.Version,
	})

	err := d.build.SaveOutput(
		plan.Type,
		source,
		resourceTypes,
		info.Version,
		db.NewResourceConfigMetadataFields(info.Metadata),
		plan.Name,
		plan.Resource,
	)
	if err != nil {
		logger.Error("failed-to-save-output", err)
		return
	}
}

func NewTaskDelegate(build db.Build, planID atc.PlanID, credVarsTracker vars.CredVarsTracker, clock clock.Clock) exec.TaskDelegate {
	return &taskDelegate{
		BuildStepDelegate: NewBuildStepDelegate(build, planID, credVarsTracker, clock),

		eventOrigin: event.Origin{ID: event.OriginID(planID)},
		build:       build,
	}
}

type taskDelegate struct {
	exec.BuildStepDelegate

	build       db.Build
	eventOrigin event.Origin
}

func (d *taskDelegate) Initializing(logger lager.Logger, taskConfig atc.TaskConfig) {
	err := d.build.SaveEvent(event.InitializeTask{
		Origin:     d.eventOrigin,
		Time:       time.Now().Unix(),
		TaskConfig: event.ShadowTaskConfig(taskConfig),
	})
	if err != nil {
		logger.Error("failed-to-save-initialize-task-event", err)
		return
	}

	logger.Info("initializing")
}

func (d *taskDelegate) Starting(logger lager.Logger, taskConfig atc.TaskConfig) {
	err := d.build.SaveEvent(event.StartTask{
		Origin:     d.eventOrigin,
		Time:       time.Now().Unix(),
		TaskConfig: event.ShadowTaskConfig(taskConfig),
	})
	if err != nil {
		logger.Error("failed-to-save-initialize-task-event", err)
		return
	}

	logger.Debug("starting")
}

func (d *taskDelegate) Finished(logger lager.Logger, exitStatus exec.ExitStatus) {
	// PR#4398: close to flush stdout and stderr
	d.Stdout().(io.Closer).Close()
	d.Stderr().(io.Closer).Close()

	err := d.build.SaveEvent(event.FinishTask{
		ExitStatus: int(exitStatus),
		Time:       time.Now().Unix(),
		Origin:     d.eventOrigin,
	})
	if err != nil {
		logger.Error("failed-to-save-finish-event", err)
		return
	}

	logger.Info("finished", lager.Data{"exit-status": exitStatus})
}

func NewCheckDelegate(check db.Check, planID atc.PlanID, credVarsTracker vars.CredVarsTracker, clock clock.Clock) exec.CheckDelegate {
	return &checkDelegate{
		BuildStepDelegate: NewBuildStepDelegate(nil, planID, credVarsTracker, clock),

		eventOrigin: event.Origin{ID: event.OriginID(planID)},
		check:       check,
		clock:       clock,
	}
}

type checkDelegate struct {
	exec.BuildStepDelegate

	check       db.Check
	eventOrigin event.Origin
	clock       clock.Clock
}

func (d *checkDelegate) SaveVersions(versions []atc.Version) error {
	return d.check.SaveVersions(versions)
}

func (*checkDelegate) Stdout() io.Writer                                 { return ioutil.Discard }
func (*checkDelegate) Stderr() io.Writer                                 { return ioutil.Discard }
func (*checkDelegate) ImageVersionDetermined(db.UsedResourceCache) error { return nil }
func (*checkDelegate) Errored(lager.Logger, string)                      { return }

func NewBuildStepDelegate(
	build db.Build,
	planID atc.PlanID,
	credVarsTracker vars.CredVarsTracker,
	clock clock.Clock,
) *buildStepDelegate {
	return &buildStepDelegate{
		build:           build,
		planID:          planID,
		clock:           clock,
		credVarsTracker: credVarsTracker,
		stdout:          nil,
		stderr:          nil,
	}
}

type buildStepDelegate struct {
	build           db.Build
	planID          atc.PlanID
	clock           clock.Clock
	credVarsTracker vars.CredVarsTracker
	stderr          io.Writer
	stdout          io.Writer
}

func (delegate *buildStepDelegate) Variables() vars.CredVarsTracker {
	return delegate.credVarsTracker
}

func (delegate *buildStepDelegate) ImageVersionDetermined(resourceCache db.UsedResourceCache) error {
	return delegate.build.SaveImageResourceVersion(resourceCache)
}

type credVarsIterator struct {
	line string
}

func (it *credVarsIterator) YieldCred(name, value string) {
	for _, lineValue := range strings.Split(value, "\n") {
		it.line = strings.Replace(it.line, lineValue, "((redacted))", -1)
	}
}

func (delegate *buildStepDelegate) buildOutputFilter(str string) string {
	it := &credVarsIterator{line: str}
	delegate.credVarsTracker.IterateInterpolatedCreds(it)
	return it.line
}

func (delegate *buildStepDelegate) Stdout() io.Writer {
	if delegate.stdout == nil {
		delegate.stdout = newDBEventWriterWithSecretRedaction(
			delegate.build,
			event.Origin{
				Source: event.OriginSourceStdout,
				ID:     event.OriginID(delegate.planID),
			},
			delegate.clock,
			delegate.buildOutputFilter,
		)
	}
	return delegate.stdout
}

func (delegate *buildStepDelegate) Stderr() io.Writer {
	if delegate.stderr == nil {
		delegate.stderr = newDBEventWriterWithSecretRedaction(
			delegate.build,
			event.Origin{
				Source: event.OriginSourceStderr,
				ID:     event.OriginID(delegate.planID),
			},
			delegate.clock,
			delegate.buildOutputFilter,
		)
	}
	return delegate.stderr
}

func (delegate *buildStepDelegate) Errored(logger lager.Logger, message string) {
	err := delegate.build.SaveEvent(event.Error{
		Message: message,
		Origin: event.Origin{
			ID: event.OriginID(delegate.planID),
		},
		Time: delegate.clock.Now().Unix(),
	})
	if err != nil {
		logger.Error("failed-to-save-error-event", err)
	}
}

func newDBEventWriter(build db.Build, origin event.Origin, clock clock.Clock) io.Writer {
	return &dbEventWriter{
		build:  build,
		origin: origin,
		clock:  clock,
	}
}

func newDBEventWriterWithSecretRedaction(build db.Build, origin event.Origin, clock clock.Clock, filter exec.BuildOutputFilter) io.WriteCloser {
	return &dbEventWriterWithSecretRedaction{
		dbEventWriter: dbEventWriter{
			build:  build,
			origin: origin,
			clock:  clock,
		},
		filter: filter,
	}
}

type dbEventWriter struct {
	build    db.Build
	origin   event.Origin
	clock    clock.Clock
	dangling []byte
}

func (writer *dbEventWriter) Write(data []byte) (int, error) {
	text := writer.writeDangling(data)
	if text == nil {
		return len(data), nil
	}

	err := writer.saveLog(string(text))
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func (writer *dbEventWriter) writeDangling(data []byte) []byte {
	text := append(writer.dangling, data...)

	checkEncoding, _ := utf8.DecodeLastRune(text)
	if checkEncoding == utf8.RuneError {
		writer.dangling = text
		return nil
	}

	writer.dangling = nil
	return text
}

func (writer *dbEventWriter) saveLog(text string) error {
	return writer.build.SaveEvent(event.Log{
		Time:    writer.clock.Now().Unix(),
		Payload: text,
		Origin:  writer.origin,
	})
}

type dbEventWriterWithSecretRedaction struct {
	dbEventWriter
	filter exec.BuildOutputFilter
}

func (writer *dbEventWriterWithSecretRedaction) Write(data []byte) (int, error) {
	var text []byte

	if data != nil {
		text = writer.writeDangling(data)
		if text == nil {
			return len(data), nil
		}
	} else {
		if writer.dangling == nil || len(writer.dangling) == 0 {
			return 0, nil
		}
		text = writer.dangling
	}

	payload := string(text)
	if data != nil {
		idx := strings.LastIndex(payload, "\n")
		if idx >= 0 && idx < len(payload) {
			// Cache content after the last new-line, and proceed contents
			// before the last new-line.
			writer.dangling = ([]byte)(payload[idx+1:])
			payload = payload[:idx+1]
		} else {
			// No new-line found, then cache the log.
			writer.dangling = text
			return len(data), nil
		}
	}

	payload = writer.filter(payload)
	err := writer.saveLog(payload)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func (writer *dbEventWriterWithSecretRedaction) Close() error {
	writer.Write(nil)
	return nil
}
