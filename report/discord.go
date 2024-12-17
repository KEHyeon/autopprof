package report

import (
	"context"
	"fmt"
	"io"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
)

const (
// reportTimeLayout = "2006-01-02T150405.MST"
// cpuCommentFmt       = ":rotating_light:[CPU] usage (*%.2f%%*) > threshold (*%.2f%%*)"
// memCommentFmt       = ":rotating_light:[MEM] usage (*%.2f%%*) > threshold (*%.2f%%*)"
// goroutineCommentFmt = ":rotating_light:[GOROUTINE] count (*%d*) > threshold (*%d*)"
)

type DiscordReporter struct {
	url     string
	appName string
	Client  webhook.Client
}

type DiscordReporterOption struct {
	url     string
	appName string
}

func DiscordSlackReporter(opt *DiscordReporterOption) (*DiscordReporter, error) {
	client, err := webhook.NewWithURL(opt.url)
	if err != nil {
		return nil, err
	}
	return &DiscordReporter{
		url:     opt.url,
		appName: opt.appName,
		Client:  client,
	}, nil
}

// ReportCPUProfile sends the CPU profiling data to the Discord
func (s *DiscordReporter) ReportCPUProfile(
	ctx context.Context, r io.Reader, ci CPUInfo,
) error {
	message := discord.NewWebhookMessageCreateBuilder().SetContent("CPU프로파일일").SetUsername("문진진").Build()
	_, err := s.Client.CreateMessage(message)
	if err != nil {
		fmt.Println("메시지 전송 실패:", err)
		return err
	}
	return nil
}

// ReportHeapProfile sends the heap profiling data to the Discord.
func (s *DiscordReporter) ReportHeapProfile(
	ctx context.Context, r io.Reader, mi MemInfo,
) error {
	message := discord.NewWebhookMessageCreateBuilder().SetContent("HEAP프로파일").SetUsername("문진진").Build()
	_, err := s.Client.CreateMessage(message)
	if err != nil {
		fmt.Println("메시지 전송 실패:", err)
		return err
	}
	return nil
}

// ReportGoroutineProfile sends the goroutine profiling data to the Discord.
func (s *DiscordReporter) ReportGoroutineProfile(
	ctx context.Context, r io.Reader, gi GoroutineInfo,
) error {
	message := discord.NewWebhookMessageCreateBuilder().SetContent("고루틴프로파일").SetUsername("문진진").Build()
	_, err := s.Client.CreateMessage(message)
	if err != nil {
		fmt.Println("메시지 전송 실패:", err)
		return err
	}
	return nil
}
