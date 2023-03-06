package channels

import (
	log "github.com/sirupsen/logrus"

	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/domain"
)

var _ IChannel = new(Writer)

type Writer struct {
	next IChannel

	logger *log.Entry

	chainName string

	cacheWriter *domain.CacheFileWriter
}

func NewWriterMW(svc IChannel, chainName string, logger *log.Logger, homeDir, dir, filename string) IChannel {

	entry := logger.WithFields(log.Fields{
		"chain_name": chainName,
	})
	cacheWriter := domain.NewCacheFileWriter(homeDir, dir, filename)
	return &Writer{
		next:        svc,
		chainName:   chainName,
		cacheWriter: cacheWriter,
		logger:      entry,
	}
}

func (w *Writer) Relay() error {
	err := w.next.Relay()
	if err != nil {
		return err
	}
	w.Context().Height()
	ctx := w.next.Context()
	w.cacheWriter.Write(ctx.Height())
	return nil
}

func (w *Writer) IsNotRelay() bool {
	return w.next.IsNotRelay()
}

func (w *Writer) Context() *domain.Context {
	return w.next.Context()
}
