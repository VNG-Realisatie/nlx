package configservice

import (
	"context"

	"go.nlx.io/nlx/config-api/configapi"
)

//ListAnnouncements returns components who have announced them selfs to the config-API
func (s *ConfigService) ListAnnouncements(ctx context.Context, req *configapi.Empty) (*configapi.ListAnnouncementsResponse, error) {
	s.logger.Info("rpc request ListAnnouncements")
	resp := &configapi.ListAnnouncementsResponse{}

	for _, value := range s.connections {
		resp.Announcements = append(resp.Announcements, &configapi.ListAnnouncementsResponse_Announcement{
			ComponentName: value.id,
		})
	}
	return resp, nil
}
