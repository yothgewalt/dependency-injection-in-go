package libraries

import (
	"fmt"
	"os"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

type EchoLogger struct {
	*Logger
}

type FxLogger struct {
	*Logger
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

func NewLogger() Logger {
	if globalLogger == nil {
		logger := initLogger(NewEnvironment())
		globalLogger = &logger
	}

	return *globalLogger
}

func (l Logger) NewEchoLogger() EchoLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return EchoLogger{Logger: initSugarLogger(logger)}
}

func (l *Logger) NewFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return &FxLogger{Logger: initSugarLogger(logger)}
}

func (f *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		f.Logger.Debug("OnStart hook failed: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			f.Logger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			f.Logger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		f.Logger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			f.Logger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			f.Logger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		f.Logger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			f.Logger.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			f.Logger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		f.Logger.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			f.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			f.Logger.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
		}
	}
}

func initSugarLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func initLogger(environment Environment) Logger {
	configuration := zap.NewDevelopmentConfig()
	logOutput := os.Getenv("LOG_OUTPUT")

	if environment.EnvironmentMode == "development" {
		fmt.Println("zapcore at encode level")
		configuration.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if environment.EnvironmentMode == "production" && logOutput != "" {
		configuration.OutputPaths = []string{logOutput}
	}

	logLevel := os.Getenv("LOG_LEVEL")
	level := zap.PanicLevel

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}

	configuration.Level.SetLevel(level)

	zapLogger, _ = configuration.Build()
	logger := initSugarLogger(zapLogger)

	return *logger
}

func (e EchoLogger) Write(p []byte) (n int, err error) {
	e.Info(string(p))

	return len(p), nil
}

func (f FxLogger) Printf(str string, args ...interface{}) {
	if len(args) > 0 {
		f.Debugf(str, args...)
	}

	f.Debug(str)
}
