// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/queue"
	"github.com/zibbp/ganymede/ent/vod"
	"github.com/zibbp/ganymede/internal/utils"
)

// Vod is the model entity for the Vod schema.
type Vod struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ExtID holds the value of the "ext_id" field.
	ExtID string `json:"ext_id,omitempty"`
	// Platform holds the value of the "platform" field.
	// The platform the VOD is from, takes an enum.
	Platform utils.VodPlatform `json:"platform,omitempty"`
	// Type holds the value of the "type" field.
	// The type of VOD, takes an enum.
	Type utils.VodType `json:"type,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// Views holds the value of the "views" field.
	Views int `json:"views,omitempty"`
	// Resolution holds the value of the "resolution" field.
	Resolution string `json:"resolution,omitempty"`
	// Processing holds the value of the "processing" field.
	// Whether the VOD is currently processing.
	Processing bool `json:"processing,omitempty"`
	// ThumbnailPath holds the value of the "thumbnail_path" field.
	ThumbnailPath string `json:"thumbnail_path,omitempty"`
	// WebThumbnailPath holds the value of the "web_thumbnail_path" field.
	WebThumbnailPath string `json:"web_thumbnail_path,omitempty"`
	// VideoPath holds the value of the "video_path" field.
	VideoPath string `json:"video_path,omitempty"`
	// ChatPath holds the value of the "chat_path" field.
	ChatPath string `json:"chat_path,omitempty"`
	// ChatVideoPath holds the value of the "chat_video_path" field.
	ChatVideoPath string `json:"chat_video_path,omitempty"`
	// InfoPath holds the value of the "info_path" field.
	InfoPath string `json:"info_path,omitempty"`
	// StreamedAt holds the value of the "streamed_at" field.
	// The time the VOD was streamed.
	StreamedAt time.Time `json:"streamed_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VodQuery when eager-loading is set.
	Edges        VodEdges `json:"edges"`
	channel_vods *uuid.UUID
}

// VodEdges holds the relations/edges for other nodes in the graph.
type VodEdges struct {
	// Channel holds the value of the channel edge.
	Channel *Channel `json:"channel,omitempty"`
	// Queue holds the value of the queue edge.
	Queue *Queue `json:"queue,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChannelOrErr returns the Channel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VodEdges) ChannelOrErr() (*Channel, error) {
	if e.loadedTypes[0] {
		if e.Channel == nil {
			// The edge channel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: channel.Label}
		}
		return e.Channel, nil
	}
	return nil, &NotLoadedError{edge: "channel"}
}

// QueueOrErr returns the Queue value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VodEdges) QueueOrErr() (*Queue, error) {
	if e.loadedTypes[1] {
		if e.Queue == nil {
			// The edge queue was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: queue.Label}
		}
		return e.Queue, nil
	}
	return nil, &NotLoadedError{edge: "queue"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vod) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case vod.FieldProcessing:
			values[i] = new(sql.NullBool)
		case vod.FieldDuration, vod.FieldViews:
			values[i] = new(sql.NullInt64)
		case vod.FieldExtID, vod.FieldPlatform, vod.FieldType, vod.FieldTitle, vod.FieldResolution, vod.FieldThumbnailPath, vod.FieldWebThumbnailPath, vod.FieldVideoPath, vod.FieldChatPath, vod.FieldChatVideoPath, vod.FieldInfoPath:
			values[i] = new(sql.NullString)
		case vod.FieldStreamedAt, vod.FieldUpdatedAt, vod.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case vod.FieldID:
			values[i] = new(uuid.UUID)
		case vod.ForeignKeys[0]: // channel_vods
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Vod", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vod fields.
func (v *Vod) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vod.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				v.ID = *value
			}
		case vod.FieldExtID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ext_id", values[i])
			} else if value.Valid {
				v.ExtID = value.String
			}
		case vod.FieldPlatform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				v.Platform = utils.VodPlatform(value.String)
			}
		case vod.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				v.Type = utils.VodType(value.String)
			}
		case vod.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case vod.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				v.Duration = int(value.Int64)
			}
		case vod.FieldViews:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field views", values[i])
			} else if value.Valid {
				v.Views = int(value.Int64)
			}
		case vod.FieldResolution:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resolution", values[i])
			} else if value.Valid {
				v.Resolution = value.String
			}
		case vod.FieldProcessing:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field processing", values[i])
			} else if value.Valid {
				v.Processing = value.Bool
			}
		case vod.FieldThumbnailPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_path", values[i])
			} else if value.Valid {
				v.ThumbnailPath = value.String
			}
		case vod.FieldWebThumbnailPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field web_thumbnail_path", values[i])
			} else if value.Valid {
				v.WebThumbnailPath = value.String
			}
		case vod.FieldVideoPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_path", values[i])
			} else if value.Valid {
				v.VideoPath = value.String
			}
		case vod.FieldChatPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chat_path", values[i])
			} else if value.Valid {
				v.ChatPath = value.String
			}
		case vod.FieldChatVideoPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chat_video_path", values[i])
			} else if value.Valid {
				v.ChatVideoPath = value.String
			}
		case vod.FieldInfoPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info_path", values[i])
			} else if value.Valid {
				v.InfoPath = value.String
			}
		case vod.FieldStreamedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field streamed_at", values[i])
			} else if value.Valid {
				v.StreamedAt = value.Time
			}
		case vod.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case vod.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case vod.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field channel_vods", values[i])
			} else if value.Valid {
				v.channel_vods = new(uuid.UUID)
				*v.channel_vods = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryChannel queries the "channel" edge of the Vod entity.
func (v *Vod) QueryChannel() *ChannelQuery {
	return (&VodClient{config: v.config}).QueryChannel(v)
}

// QueryQueue queries the "queue" edge of the Vod entity.
func (v *Vod) QueryQueue() *QueueQuery {
	return (&VodClient{config: v.config}).QueryQueue(v)
}

// Update returns a builder for updating this Vod.
// Note that you need to call Vod.Unwrap() before calling this method if this Vod
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vod) Update() *VodUpdateOne {
	return (&VodClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Vod entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vod) Unwrap() *Vod {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vod is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vod) String() string {
	var builder strings.Builder
	builder.WriteString("Vod(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", ext_id=")
	builder.WriteString(v.ExtID)
	builder.WriteString(", platform=")
	builder.WriteString(fmt.Sprintf("%v", v.Platform))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", v.Type))
	builder.WriteString(", title=")
	builder.WriteString(v.Title)
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", v.Duration))
	builder.WriteString(", views=")
	builder.WriteString(fmt.Sprintf("%v", v.Views))
	builder.WriteString(", resolution=")
	builder.WriteString(v.Resolution)
	builder.WriteString(", processing=")
	builder.WriteString(fmt.Sprintf("%v", v.Processing))
	builder.WriteString(", thumbnail_path=")
	builder.WriteString(v.ThumbnailPath)
	builder.WriteString(", web_thumbnail_path=")
	builder.WriteString(v.WebThumbnailPath)
	builder.WriteString(", video_path=")
	builder.WriteString(v.VideoPath)
	builder.WriteString(", chat_path=")
	builder.WriteString(v.ChatPath)
	builder.WriteString(", chat_video_path=")
	builder.WriteString(v.ChatVideoPath)
	builder.WriteString(", info_path=")
	builder.WriteString(v.InfoPath)
	builder.WriteString(", streamed_at=")
	builder.WriteString(v.StreamedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Vods is a parsable slice of Vod.
type Vods []*Vod

func (v Vods) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
