# Data Infrastructure Engineering Test
Our marketing team would like to understand, how often the word `taxi` is being referenced in wikipedia articles. In order to enable this capability, we need to make the articles available and analyzable internally. As an engineer in the data platform team, you are being asked to design and implement this process. You may choose any technology to complete the task at hand, but need to be able to justify your approach.

## Task Details
1. Consume wikipedia api and crawl around 10 articles (it does not matter which ones) and store them in a database of your choice.
2. Write an application that summarizes how often a keyword appeared in all articles stored this way, f.e. `taxi`.
3. Combine your build and execution in a docker environment for both the code and the database f.e `docker-compose`. 

During the next interview, we would like to discuss the technologies and design you have chosen.


## Conditions
* You may use any programming language for your new service. Feel free to pick any language you feel comfortable with and is matching the usecase.
* Please combine your results in a Dockerfile.
* Spend as much time, as you would like to on the test. We do not expect you to spend your entire weekend on this.
* Track your progress using git. No need to push to a remote repository, but we would love to see a local git history.

## Helpful Links
* https://www.mediawiki.org/wiki/API:Main_page
* https://en.wikipedia.org/w/api.php?action=query&prop=revisions&rvprop=content&rvsection=0&titles=Free_Now_(service)&format=json

