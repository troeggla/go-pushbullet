package pushbullet

import (
    "encoding/json"
    "errors"
    "net/http"
)

type channelResponse struct {
    Channels    []*Channel
}

// A Channel is a PushBullet channel
type Channel struct {
    Iden         string  `json:"iden"`
    Name         string  `json:"name"`
    Tag          string  `json:"tag"`
    Description  string  `json:"description"`
    Active       bool    `json:"active"`
    Created      float32 `json:"created"`
    Modified     float32 `json:"modified"`
    ImageUrl     string  `json:"image_url"`
}

// Channels fetches a list of channels from PushBullet.
func (c *Client) Channels() ([]*Channel, error) {
    req := c.buildRequest("/channels", nil)
    resp, err := c.Client.Do(req)

    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        var errjson errorResponse

        dec := json.NewDecoder(resp.Body)
        err = dec.Decode(&errjson)

        if err == nil {
            return nil, &errjson.ErrResponse
        }

        return nil, errors.New(resp.Status)
    }

    var chResp channelResponse
    dec := json.NewDecoder(resp.Body)
    err = dec.Decode(&chResp)

    if err != nil {
        return nil, err
    }

    devices := append(chResp.Channels)
    return devices, nil
}
