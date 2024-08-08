// package cmd

// import (
//     "fmt"
//     "os"
//     "os/exec"
//     "strings"

//     "github.com/spf13/cobra"
// )

// var videoURL string

// // downloadCmd represents the download command
// var downloadCmd = &cobra.Command{
//     Use:   "download",
//     Short: "Download Instagram video",
//     Long:  `Download video from Instagram using the provided URL.`,
//     Run: func(cmd *cobra.Command, args []string) {
//         downloadVideo(videoURL)
//     },
// }

// func init() {
//     rootCmd.AddCommand(downloadCmd)

//     downloadCmd.Flags().StringVarP(&videoURL, "url", "u", "", "URL of the Instagram video")
//     downloadCmd.MarkFlagRequired("url")
// }

// func downloadVideo(url string) {
//     // Check if URL is valid
//     if !strings.Contains(url, "instagram.com") {
//         fmt.Println("Invalid Instagram URL")
//         return
//     }

//     // Prepare the command to download the video using yt-dlp
//     cmd := exec.Command("yt-dlp", "--no-check-certificate", "-o", "video.mp4", url)

//     // Set the output to the console
//     cmd.Stdout = os.Stdout
//     cmd.Stderr = os.Stderr

//     // Run the command
//     err := cmd.Run()
//     if err != nil {
//         fmt.Println("Error downloading video:", err)
//         return
//     }

//     fmt.Println("Video downloaded successfully!")
// }




///-----------------kode ini untuk mendownload semua jenis platform--------------------



package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "github.com/spf13/cobra"
)

var urls []string

var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download videos from any supported platform",
    Run: func(cmd *cobra.Command, args []string) {
        printCurrentDirectory()
        for _, url := range urls {
            downloadVideo(url)
        }
    },
}

func init() {
    rootCmd.AddCommand(downloadCmd)
    downloadCmd.Flags().StringArrayVarP(&urls, "url", "u", []string{}, "URLs of the videos to download (comma-separated)")
    downloadCmd.MarkFlagRequired("url")
}

func printCurrentDirectory() {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return
    }
    fmt.Println("Current working directory:", cwd)
}

func downloadVideo(url string) {
    outputFileName := fmt.Sprintf("video_%s.mp4", strings.ReplaceAll(url, "/", "_"))
    cmd := exec.Command("yt-dlp", "--no-check-certificate", "-o", outputFileName, url)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error downloading video:", err)
        return
    }
    fmt.Printf("Video downloaded successfully: %s\n", outputFileName)
}
