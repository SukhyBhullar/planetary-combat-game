package main

import (
	"context"
	"fmt"

	pb "github.com/SukhyBhullar/planetary-combat-game/planetary"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	sub := rdb.Subscribe(ctx, "messages")
	ch := sub.Channel()
	fmt.Println("Connected")
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		data := []byte(msg.Payload)
		game := &pb.Game{}
		proto.Unmarshal(data, game)
		if err := proto.Unmarshal(data, &pb.Game{}); err != nil {
			fmt.Println("Failed to parse address book:", err)
		}
		fmt.Printf("Char: %s\n", *game.Charname)

	}

	// err := rdb.Subscribe(ctx, "messages").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist
}
