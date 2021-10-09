package main

func FillPost(Newpost Post)(Post) {
	var post Post
	if Newpost.Caption != "" {
		post.Caption = Newpost.Caption
	}
	if Newpost.ImgURL != "" {
		post.ImgURL = Newpost.ImgURL
	}
	if Newpost.StartTime != "" {
		post.StartTime = Newpost.StartTime
	}
	if Newpost.EndTime != "" {
		post.EndTime = Newpost.EndTime
	}
	return post
}