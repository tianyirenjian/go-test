package main
import (
    "fmt"
    "sort"
)

type tweet struct {
    Id int
    Uid int
    Time int
}

type sortTweet []tweet

func (st sortTweet) Len() int {
    return len(st)
}

func (st sortTweet) Swap(i, j int) {
    st[i], st[j] = st[j], st[i]
}

func (st sortTweet) Less(i, j int) bool {
    return st[i].Time > st[j].Time
}

type Twitter struct {
    twitters []tweet
    time int
    follow map[int][]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
    tw := Twitter{}
    tw.follow = make(map[int][]int)
    return tw
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.twitters = append(this.twitters, tweet{tweetId, userId, this.time})
    this.time++
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    ids := []int{userId}
    followed := this.follow[userId]
    ids = append(ids, followed...)
    tweets := sortTweet{}
    for _, uid := range ids {
        for _,tw := range this.twitters {
            if tw.Uid == uid && !arrayHas(tweets, tw.Id) {
                tweets = append(tweets, tw)
            }
        }
    }
    sort.Sort(tweets)
    if len(tweets) > 10 {
        tweets = tweets[:10]
    }
    ret := []int{}
    for _, t := range tweets {
        ret = append(ret, t.Id)
    }
    return ret
}

func arrayHas(arr sortTweet, find int) bool {
    for _, tweet := range arr {
        if tweet.Id == find {
            return true
        }
    }
    return false
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    followed := this.follow[followerId]
    // 防止重复关注
    has := false
    for _,v := range followed {
        if v == followeeId {
            has = true
            break
        }
    }
    if !has {
        followed = append(followed, followeeId)
    }
    this.follow[followerId] = followed
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    followed := this.follow[followerId]
    for i, v := range followed {
        if v == followeeId {
            followed = append(followed[:i], followed[i+1:]...)
        }
    }
    this.follow[followerId] = followed
}


func main() {
    obj := Constructor()
    obj.PostTweet(1, 5)
    obj.Follow(1, 1)
    param_2 := obj.GetNewsFeed(1)
    fmt.Println(obj)
    fmt.Println(param_2)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
