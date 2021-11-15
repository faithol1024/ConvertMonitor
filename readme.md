# What is this?
This Simple script is used to convert Datadog Dashboard to NewRelic.
<br>This script is build with specific dashboard layout in mind, so it may not work for everyone

## How To Use
- Copy your datadog dashboard json from the topright menu <br> ![image](https://user-images.githubusercontent.com/46881265/141757033-fb8c8682-1e68-4e35-9d20-5ca92cefc7af.png)
- paste the json into datadog.json
- use command `make gorun`
- The resulting output will be printed on `newrelic.json`
- on your newrelic dashboard page, click `import dashboard` ![image](https://user-images.githubusercontent.com/46881265/141762667-30c59d91-764a-49aa-8aaa-37acfbb8caaa.png)
- voila! your dashboard will be created automatically

## Side notes
- newrelic custom metric dashboard has maximum of 100 widgets. this script will make new page if the limit already reached
