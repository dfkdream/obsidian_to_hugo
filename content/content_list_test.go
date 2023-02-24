package content

import (
	"obsidian_md/config"
	"os"
	"testing"
	"time"
)

func TestFromDirectory(t *testing.T) {
	testSiteConfig, _ := config.FromFile("../test_site/config.yml")
	loc, _ := time.LoadLocation("Asia/Seoul")
	fTime := time.Date(2023, time.February, 23, 23, 16, 25, 822648616, loc)
	_ = os.Chtimes("../test_site/content/posts2/post5.md", fTime, fTime)

	type contentTest struct {
		ObsidianIdentifier string
		HugoIdentifier     string
	}

	type args struct {
		path   string
		config config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    []contentTest
		wantErr bool
	}{
		{
			name: "test site",
			args: args{
				path:   "../test_site/content",
				config: testSiteConfig,
			},
			want: []contentTest{
				{
					ObsidianIdentifier: "outside_posts",
					HugoIdentifier:     "/outside_posts/",
				},
				{
					ObsidianIdentifier: "posts/Pasted image 20230222161509.png",
					HugoIdentifier:     "/posts/Pasted image 20230222161509.png",
				},
				{
					ObsidianIdentifier: "posts/post1",
					HugoIdentifier:     "/2023/02/22/post1/",
				},
				{
					ObsidianIdentifier: "posts/post2",
					HugoIdentifier:     "/2023/02/23/post2/",
				},
				{
					ObsidianIdentifier: "posts/post3",
					HugoIdentifier:     "/2023/02/24/post3/",
				},
				{
					ObsidianIdentifier: "posts/post_with_same_filename",
					HugoIdentifier:     "/2023/02/25/post_with_same_filename/",
				},
				{
					ObsidianIdentifier: "posts/subdirectory/post_with_same_filename",
					HugoIdentifier:     "/2023/02/26/post_with_same_filename/",
				},
				{
					ObsidianIdentifier: "posts/subdirectory2/post4",
					HugoIdentifier:     "/2023/02/24/post4/",
				},
				{
					ObsidianIdentifier: "posts2/post5",
					HugoIdentifier:     "/posts2/post5/",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromDirectory(tt.args.path, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for idx, v := range tt.want {
				if v.HugoIdentifier != got[idx].HugoIdentifier() || v.ObsidianIdentifier != got[idx].ObsidianIdentifier() {
					t.Errorf("FromDirectory() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
