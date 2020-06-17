package handler

import (
	"context"

	"github.com/gofiber/fiber"

	"github.com/escaletech/tog-go/sessions"
)

type sessionCreator interface {
	Session(context.Context, string, string, *sessions.SessionOptions) sessions.Session
}

type handler struct {
	client sessionCreator
}

func (h *handler) Session(c *fiber.Ctx) {
	namespace, sessionID := c.Params("namespace"), c.Params("session")

	opt := parseOptions(c)
	session := h.client.Session(c.Context(), namespace, sessionID, &opt)

	c.Status(fiber.StatusOK).JSON(SessionResponse{
		Namespace: namespace,
		ID:        sessionID,
		Flags:     session,
	})
}

func parseOptions(c *fiber.Ctx) sessions.SessionOptions {
	args := c.Fasthttp.QueryArgs()

	force := sessions.Session{}
	for _, f := range args.PeekMulti("enable") {
		force[string(f)] = true
	}
	for _, f := range args.PeekMulti("disable") {
		force[string(f)] = false
	}
	if len(force) == 0 {
		force = nil
	}

	tq := args.PeekMulti("traits")
	var traits []string
	if len(tq) > 0 {
		traits = make([]string, len(tq))
		for i, t := range tq {
			traits[i] = string(t)
		}
	}

	return sessions.SessionOptions{
		Traits: traits,
		Force:  force,
	}
}

type SessionResponse struct {
	Namespace string           `json:"namespace"`
	ID        string           `json:"id"`
	Flags     sessions.Session `json:"flags"`
}
