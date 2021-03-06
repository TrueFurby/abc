package main

/* Writer: interface for storing the app data*/
type Writer interface {
	getBoards(apiKey string) (currBoards []boards, err error)
	getActiveThreadsForBoard(apiKey string, boardID int) (activeThreads []threads, err error)
	getPostsForThread(apiKey string, threadID int) (currPosts []threadPosts, err error)
	addPostToThread(threadID int, threadBodyPost string, attachmentURL *string, clientRemoteAddr string) (err error)
	addThread(boardID int, threadName string) (threads, error)
	isThreadLimitReached(boardID int) (bool, error)
	isPostLimitReached(threadID int) (bool, threads, error)
	archiveThread(threadID int)
}
