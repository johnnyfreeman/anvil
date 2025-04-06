package progress

import (
	"github.com/charmbracelet/log"
)

type LogReporter struct {
	logger *log.Logger
}

func NewLogReporter() *LogReporter {
	return &LogReporter{logger: log.Default()}
}

func (r *LogReporter) Start(task string) {
	r.logger.Infof("🚀 Starting: %s", task)
}

func (r *LogReporter) Success(task string) {
	r.logger.Infof("✅ Success: %s", task)
}

func (r *LogReporter) Failure(task string, err error) {
	r.logger.Errorf("❌ Failed: %s (%v)", task, err)
}

func (r *LogReporter) Done() {
	r.logger.Info("🎉 All tasks complete")
}
