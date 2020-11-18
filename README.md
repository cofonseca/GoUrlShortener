# Reb.red - A Simple URL Shortener
**[Rebred](https://reb.red)** (_to breed [a url] for a second or subsequent time_) is a URL shortener and proxy, similar to bit.ly, written mostly in Go.

Paste in a long and ugly link, and Rebred will generate a shorter and much friendlier link for you. You can optionally choose your own nickname/sub-path for your link. Your new link will last forever.

The purpose of this app was primarily to help me learn Go programming, and refresh my memory on how to write a simple full-stack web application.

Rebred was written using the Go standard libraries, and leverages Cloudflare and Google Cloud products in order to function. The application is built and packaged into a container using Google Cloud Build, is stored in Google Container Registry (GCR), runs in Google Cloud Run, and uses Google Firestore as its database. 

Since this is a personal project, I am not looking for any community contributions at this time, but please feel free to submit an issue to report bugs or request a feature. Thanks for looking!