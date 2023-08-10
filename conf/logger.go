package conf

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	core := zapcore.NewCore(getEncoder())

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getwriteSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync()
}
