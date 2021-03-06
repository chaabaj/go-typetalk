package v2

import (
	"context"
	"time"

	. "github.com/nulab/go-typetalk/typetalk/internal"
	. "github.com/nulab/go-typetalk/typetalk/shared"
)

type LikesService service

type Post struct {
	ID            int               `json:"id"`
	TopicID       int               `json:"topicId"`
	Topic         Topic             `json:"topic"`
	ReplyTo       int               `json:"replyTo"`
	Message       string            `json:"message"`
	Account       Account           `json:"account"`
	Attachments   []*AttachmentFile `json:"attachments"`
	Links         []interface{}     `json:"links"`
	DirectMessage *DirectMessage    `json:"directMessage"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

type Account struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	FullName   string     `json:"fullName"`
	Suggestion string     `json:"suggestion"`
	ImageURL   string     `json:"imageUrl"`
	IsBot      bool       `json:"isBot"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type Like struct {
	ID        int        `json:"id"`
	PostID    int        `json:"postId"`
	TopicID   int        `json:"topicId"`
	Comment   string     `json:"comment"`
	Account   *Account   `json:"account"`
	CreatedAt *time.Time `json:"createdAt"`
}

type AttachmentFile struct {
	ContentType string `json:"contentType"`
	FileKey     string `json:"fileKey"`
	FileName    string `json:"fileName"`
	FileSize    int    `json:"fileSize"`
}

type DirectMessage struct {
	Account *Account `json:"account"`
	Status  *Status  `json:"status"`
}

type Status struct {
	Presence *string     `json:"presence"`
	Web      interface{} `json:"web"`
	Mobile   interface{} `json:"mobile"`
}

type LikedPost struct {
	Post          *Post          `json:"post"`
	Likes         []*Like        `json:"likes"`
	DirectMessage *DirectMessage `json:"directMessage"`
}

type DiscoverLikedPost struct {
	*LikedPost
}

type ReceiveLikedPost struct {
	*LikedPost
}

type GiveLikedPost struct {
	*ReceiveLikedPost
	MyLike *MyLike `json:"myLike"`
}

type MyLike struct {
	ID        int       `json:"id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetLikesOptions struct {
	From     int    `json:"from,omitempty"`
}

type getLikesOptions struct {
	*GetLikesOptions
	SpaceKey string `json:"spaceKey"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-likes-receive/
func (s *LikesService) GetLikesReceive(ctx context.Context, spaceKey string, opt *GetLikesOptions) ([]*ReceiveLikedPost, *Response, error) {
	u, err := AddQueries("likes/receive", &getLikesOptions{GetLikesOptions: opt, SpaceKey: spaceKey})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		LikedPosts []*ReceiveLikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.LikedPosts, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-likes-give/
func (s *LikesService) GetLikesGive(ctx context.Context, spaceKey string, opt *GetLikesOptions) ([]*GiveLikedPost, *Response, error) {
	u, err := AddQueries("likes/give", &getLikesOptions{GetLikesOptions: opt, SpaceKey: spaceKey})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		GiveLikedPost []*GiveLikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.GiveLikedPost, resp, nil
	}
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/get-likes-discover/
func (s *LikesService) GetLikesDiscover(ctx context.Context, spaceKey string, opt *GetLikesOptions) ([]*DiscoverLikedPost, *Response, error) {
	u, err := AddQueries("likes/discover", &getLikesOptions{GetLikesOptions: opt, SpaceKey: spaceKey})
	if err != nil {
		return nil, nil, err
	}
	var result *struct {
		DiscoverLikedPost []*DiscoverLikedPost `json:"likedPosts"`
	}
	if resp, err := s.client.Get(ctx, u, &result); err != nil {
		return nil, resp, err
	} else {
		return result.DiscoverLikedPost, resp, nil
	}
}

type ReadReceivedLikesOptions struct {
	LikeId int `json:"likeId,omitempty"`
}

type readReceivedLikesOptions struct {
	*ReadReceivedLikesOptions
	SpaceKey string `json:"spaceKey"`
}

type ReadReceivedLikesResult struct {
	Like struct {
		Receive struct {
			HasUnread  bool `json:"hasUnread"`
			ReadLikeID int  `json:"readLikeId"`
		} `json:"receive"`
	} `json:"like"`
}

// Typetalk API docs: https://developer.nulab-inc.com/docs/typetalk/api/2/save-read-likes/
func (s *LikesService) ReadReceivedLikes(ctx context.Context, spaceKey string, opt *ReadReceivedLikesOptions) (*ReadReceivedLikesResult, *Response, error) {
	u := "likes/receive/bookmark/save"
	var result *ReadReceivedLikesResult
	if resp, err := s.client.Post(ctx, u, &readReceivedLikesOptions{ReadReceivedLikesOptions: opt, SpaceKey: spaceKey}, result); err != nil {
		return nil, resp, err
	} else {
		return result, resp, nil
	}
}
