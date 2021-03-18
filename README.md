#My Personal Site
I made this site to be a landing page for potential clients/employers. 
After trying to create a design multiple times, I accepted something I 
mostly already knew, that I'm not a designer. I decided to use a free 
design by HTML5UP. I did not use the code from this design, rather I
used it as a visual aid while creating the site. 

As of now it is using HTML, CSS and Go. I enjoy using Go as a backend
as it is a great language to work with and very effecient

###Hosting
The site is hosted on an AWS EC2 instance. I have a GitHub Action set up
to trigger CodeDeploy each time I push to the repository. The config file
kills the existing Go server if it is running and removes the current site
directory. After adding a directory with the new files, it restarts the
server.