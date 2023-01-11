package main

import (
	"context"
	"fmt"
)

type Publisher struct {
	name        string
	ctx         context.Context
	subscribeCh chan chan<- string // 구독자 추가 여부 알기 위한 채널
	publishCh   chan string        // 메시지 발행을 위한 채널
	subscribers []chan<- string    // 구독자 목록
}

func NewPublisher(name string, ctx context.Context) *Publisher {
	return &Publisher{
		name:        name,
		ctx:         ctx,
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

// 구독자 추가
func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscribeCh <- sub // 구독자가 chan<- string 타입의 msg를 넣어줄 것임
}

// 메시지 발행
func (p *Publisher) Publish(msg string) {
	p.publishCh <- fmt.Sprintf("%s, from %s", msg, p.name)
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh: // 구독자가 생기면 구독자 목록에 추가
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh: // 발행할 메시지가 생기면 구독자들에게 발행함
			for _, subscriber := range p.subscribers {
				subscriber <- msg // 구독자의 msgCh에 msg 보냄
			}
		case <-p.ctx.Done(): // 종료
			wg.Done()
			return
		}
	}
}
