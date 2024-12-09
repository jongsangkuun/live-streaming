package main

import (
	"fmt"

	"io"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/yutopp/go-rtmp"
	"github.com/yutopp/go-rtmp/message"
)

type Handler struct{}

func (h *Handler) OnConnect(timestamp uint32, cmd *message.NetConnectionConnect) error {
	log.Info("OnConnect")
	return nil
}

func (h *Handler) OnCreateStream(timestamp uint32, cmd *message.NetConnectionCreateStream) error {
	log.Info("OnCreateStream")
	return nil
}

func (h *Handler) OnReleaseStream(timestamp uint32, cmd *message.NetConnectionReleaseStream) error {
	log.Info("OnReleaseStream")
	return nil
}

func (h *Handler) OnDeleteStream(timestamp uint32, cmd *message.NetStreamDeleteStream) error {
	log.Info("OnDeleteStream")
	return nil
}

func (h *Handler) OnPublish(ctx *rtmp.StreamContext, timestamp uint32, cmd *message.NetStreamPublish) error {
	streamKey := cmd.PublishingName // 스트림 키 값을 가져옵니다.

	if isValidStreamKey(streamKey) { // 스트림 키 유효성 검사
		log.Infof("스트림 시작: %s", streamKey)
		return nil // 스트리밍 허용
	} else {
		log.Warnf("잘못된 스트림 키: %s", streamKey)
		return fmt.Errorf("잘못된 스트림 키") // 스트리밍 거부
	}
}

func isValidStreamKey(key string) bool {
	return key == "mysecretkey" // 스트림 키가 "mysecretkey"인지 확인합니다.
}

func (h *Handler) OnPlay(ctx *rtmp.StreamContext, timestamp uint32, cmd *message.NetStreamPlay) error {
	log.Info("OnPlay")
	return nil
}

func (h *Handler) OnFCPublish(timestamp uint32, cmd *message.NetStreamFCPublish) error {
	log.Info("OnFCPublish")
	return nil
}

func (h *Handler) OnFCUnpublish(timestamp uint32, cmd *message.NetStreamFCUnpublish) error {
	log.Info("OnFCUnpublish")
	return nil
}

func (h *Handler) OnSetDataFrame(timestamp uint32, data *message.NetStreamSetDataFrame) error {
	log.Info("OnSetDataFrame")
	return nil
}

func (h *Handler) OnUnknownMessage(timestamp uint32, msg message.Message) error {
	//TODO implement me
	log.Info("ON UNKNOWN MESSAGE")
	return nil
}

func (h *Handler) OnUnknownCommandMessage(timestamp uint32, cmd *message.CommandMessage) error {
	//TODO implement me
	log.Info("OnUnknownCommandMessage")
	return nil
}

func (h *Handler) OnUnknownDataMessage(timestamp uint32, data *message.DataMessage) error {
	//TODO implement me
	log.Info("OnUnknownDataMessage")
	return nil
}

func (h *Handler) OnClose() {
	//TODO implement me
	log.Info("OnClose")
}

func (h *Handler) OnServe(conn *rtmp.Conn) {
	log.Info("Stream connected")
}

func (h *Handler) OnReceiveAudio(conn *rtmp.Conn, timestamp uint32, payload io.Reader) error {
	log.Info("Audio received")
	return nil
}

func (h *Handler) OnReceiveVideo(conn *rtmp.Conn, timestamp uint32, payload io.Reader) error {
	log.Info("Video received")
	return nil
}

func (h *Handler) OnAudio(timestamp uint32, payload io.Reader) error {
	log.Info("Audio received")
	return nil
}

func (h *Handler) OnVideo(timestamp uint32, payload io.Reader) error {
	log.Info("Video received")
	return nil
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:1935")
	if err != nil {
		log.Panicf("Failed: %+v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Panicf("Failed: %+v", err)
	}

	srv := rtmp.NewServer(&rtmp.ServerConfig{
		OnConnect: func(conn net.Conn) (io.ReadWriteCloser, *rtmp.ConnConfig) {
			l := log.StandardLogger()
			//l.SetLevel(log.DebugLevel)

			h := &Handler{}

			return conn, &rtmp.ConnConfig{
				Handler: h,

				ControlState: rtmp.StreamControlStateConfig{
					DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
				},

				Logger: l,
			}
		},
	})
	if err := srv.Serve(listener); err != nil {
		log.Panicf("Failed: %+v", err)
	}
}
