---
title: "Upload Data"
date: 2022-08-19
lastmod: 2022-08-23
draft: true
---
Use this page to test data protection controls. Activities that can be tested here include form post, upload and even download(From the browse section).

## Form Post Test
This form will post data to the flonkerton api and write it to a file using the given name.

### {{< textDataForm action="/api/formpost"  >}}
<br><br>

## File Upload Test
Select a single or multiple files to upload to flonkerton. If you controls are working these files will not appear below.  

### {{< fileUploadForm action="/api/fileupload"  >}}

  
## [Uploaded Files](/uploads)
{{< iframe "/uploads" >}}