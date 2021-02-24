package postNats

import (
	"microServiceBoilerplate/configs"
	natsPb "microServiceBoilerplate/proto/generated/nats"
	pb "microServiceBoilerplate/proto/generated/post"
	"microServiceBoilerplate/services/post/instances"

	"github.com/mreza0100/golog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func newPublishers(nc *nats.Conn, lgr *golog.Core) instances.Publishers {
	return &publishers{
		lgr: lgr.With("In publishers->"),
		nc:  nc,
	}
}

type publishers struct {
	lgr *golog.Core
	nc  *nats.Conn
}

func (p *publishers) NewPost(userId, postId uint64) error {
	subject := configs.Nats.Subjects.NewPost

	{
		data := &natsPb.NewPost_EVENT{
			UserId: userId,
			PostId: postId,
		}
		msgByte, err := proto.Marshal(data)
		if err != nil {
			p.lgr.Debug.RedLog("proto.Marshal has been returning error")
			p.lgr.Debug.RedLog("Error: ", err)
			return err
		}

		err = p.nc.Publish(subject, msgByte)
		if err != nil {
			p.lgr.RedLog("in NewPost: can't publish msg")
			p.lgr.RedLog("error: ", err)
			return err
		}
	}
	return nil
}

func (p *publishers) GetUsers(userIds []uint64) (map[uint64]*pb.User, error) {
	subject := configs.Nats.Subjects.GetUsersByIds
	dbug, success := p.lgr.DebugPKG("GetProfiles", false)

	{
		var (
			request  []byte
			rawUsers *natsPb.GetUsers_REQUESTResponse
			result   map[uint64]*pb.User

			err error
		)

		{
			rawUsers = &natsPb.GetUsers_REQUESTResponse{}
		}
		{
			request, err = proto.Marshal(&natsPb.GetUsers_REQUESTRequest{
				UserIds: userIds,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil, err
			}
		}
		{
			var res *nats.Msg
			res, err := p.nc.Request(subject, request, configs.Nats.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil, err
			}
			err = proto.Unmarshal(res.Data, rawUsers)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil, err
			}
		}
		{
			result = make(map[uint64]*pb.User, len(rawUsers.UsersData))
			for _, u := range rawUsers.UsersData {
				result[u.Id] = &pb.User{
					Name:   u.Name,
					Email:  u.Email,
					Id:     u.Id,
					Gender: u.Gender,
				}
			}
		}
		success("profiles: ", result)
		return result, nil
	}
}

func (p *publishers) IsFollowingGroup(userId uint64, followings []uint64) map[uint64]bool {
	subject := configs.Nats.Subjects.IsFollowingGroup
	dbug, success := p.lgr.DebugPKG("IsFollowings", false)

	{
		var (
			response     *natsPb.IsFollowingGroup_REQUESTResponse
			byteRequest  []byte
			byteResponse []byte

			err error
		)

		{
			response = &natsPb.IsFollowingGroup_REQUESTResponse{}
		}
		{
			byteRequest, err = proto.Marshal(&natsPb.IsFollowingGroup_REQUESTRequest{
				Follower:   userId,
				Followings: followings,
			})
			if dbug("proto.Marshal")(err) != nil {
				return nil
			}
		}
		{
			res, err := p.nc.Request(subject, byteRequest, configs.Nats.Timeout)
			if dbug("p.nc.Request")(err) != nil {
				return nil
			}
			byteResponse = res.Data
		}
		{
			err = proto.Unmarshal(byteResponse, response)
			if dbug("proto.Unmarshal")(err) != nil {
				return nil
			}
		}
		{
			success(response.Result)
			return response.Result
		}
	}
}
