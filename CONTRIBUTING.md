# Contributing

We – [VNG realisatie](https://www.vngrealisatie.nl/) and the maintainers of this project – know we can only build NLX together with you. Thus we appreciate your input, enjoy feedback and welcome improvements to this project and are very open to collaboration.

We love issues and merge requests from everyone.

## 1. Problems, suggestions and questions in Issues

You don't need to change any of our code or documentation to be a contributor. Please help development by reporting problems, suggesting changes and asking questions. To do this, you can [create an issue using GitLab](https://docs.gitlab.com/ee/user/project/issues/create_new_issue.html) for this project in the [GitLab Issues for NLX](https://gitlab.com/commonground/nlx/nlx/issues).

## 2. Documentation and code in Merge Requests

If you want to add to the documentation or code of one of our projects you should push a branch and make a Merge Request. If you have never used GitLab before, get up to speed by reading about the [GitLab Flow](https://about.gitlab.com/2014/09/29/gitlab-flow/).

### 2.1. Make your changes

#### 2.1.1. Use OneFlow
This project uses the **OneFlow branching model** and workflow. When you've forked this repository, please make sure to create a feature branch following the OneFlow model. Read this [blogpost](http://endoflineblog.com/oneflow-a-git-branching-model-and-workflow) when you're not yet familiar with OneFlow.

#### 2.1.2. Add docs and tests
If you are adding code, make sure you've added and updated the relevant documentation and tests before you submit your Merge Request. Make sure to write unit tests that show the behaviour of the newly added or changed code.

### 2.2. Commit messages

#### 2.2.1. Explain your contributions
Add your changes in commits [with a message that explains them](https://robots.thoughtbot.com/5-useful-tips-for-a-better-commit-message). Document choices or decisions you make in the commit message, this will enable everyone to be informed of your choices in the future.

#### 2.2.2. Semantic Release
This project uses [semantic-release](https://semantic-release.gitbook.io/semantic-release/). When merging a MR to master, this will automatically generate our [CHANGELOG](./CHANGELOG.md) based on the commit messages and a version tag will be added.

#### 2.2.3. Conventions for commit messages
We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. This is enforced with a linter in the build pipeline. This convention requires you to pas a type and an optional scope as the commit message. The scope is based on the applications in the repository. If you are not sure which scope to use please leave the scope empty.

The type must be one of the following:

- **build**: Changes that affect the build system or external dependencies
- **ci**: Changes to our CI configuration files and scripts
- **docs**: Documentation only changes
- **feat**: A new feature
- **fix**: A bug fix
- **perf**: A code change that improves performance
- **refactor**: A code change that neither fixes a bug nor adds a feature
- **revert**: Changes that revert other changes
- **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- **test**: Adding missing tests or correcting existing tests

The available scopes are:

- auth-service
- ca-certportal
- ca-cfssl-unsafe
- common
- directory
- docs
- helm
- insight
- inway
- management
- outway
- txlog-db

### 2.3. Merge Request

#### 2.3.1. Always refer to an issue
Before starting a Merge Request, make sure there is a User Story describing what you want to achieve with the MR. [Create a story by submitting a new issue](https://gitlab.com/commonground/nlx/nlx/issues) if there is none. New issues come with a User Story template. This template helps you think from the user perspective: 'who wants this new feature and why?'

#### 2.3.2. Describe the MR

When submitting the Merge Request, please accompany it with a short description and eventually some clarifying remarks for the reviewer. Make sure the related issues are linked to the MR.

#### 2.3.3. Combine frontend and backend work in one MR
When working on a feature which requires specific capabilities of multiple developers (eg. both specialistic frontend & backend work), work as a team to make sure the MR contains a full feature instead of separate MR's per developer. By doing so, the reviewer can consider the complete solution and give more insightful feedback.

### 2.4. Improve

#### 2.4.1. Reviews
All contributions have to be reviewed. The reviewer will look at the submitted code to assure quality. The reviewer will consider at least the following:
* Does the MR meet the intended acceptance criteria?
* Are there any typos?
* Does the code meet the teams overall quality standards?
* Are there logical errors in the code?
* Is the chosen technical solution reasonably efficient?
* Are all the necessary (unit) tests added?


#### 2.4.2. Definition of Done

With MR's we make User Stories become reality.
This section describes the Definition of Done expressed in the stages a story goes through from cradle to cradle.

**1. Definition of Ready**<BR>
*Responsible: PO/DEV*<BR>
- A story passes the INVEST criteria:<BR>
“I” ndependent (of all others)<BR>
“N” egotiable (not a specific contract for features)<BR>
“V” aluable (or vertical)<BR>
“E” stimatable (to a good approximation)<BR>
“S” mall (so as to fit within an iteration)<BR>
“T” estable (in principle, even if there isn’t a test for it yet)
  - Test scenario's are written down
  - acceptance criteria are written down in a way no misinterpretations are possible
- We describe the story with the end in mind. So part of the story is the way we are going to demonstrate it to the stakeholder
- We describe the proposed solution which has been discussed during refinement

**2. From Doing to Review**<BR>
*Responsible: DEV*<BR>
- Code builds
  - linting errors are fixed
  - security issues are fixed
- Tests are added
- The developer has checked all acceptance criteria and test scenarios first
- The developer is responsible for arranging the review of the story
  - The developer will notify the team who's review his story
- Developer moves story to review

**3. From Review to Accept**<BR>
*Responsible: DEV*<BR>
- Test are in place and understood
- Solution has been checked and discussed with responsible developer
- Reviewer and developer agree upon the solution
- Reviewer marks the story as being reviewed
- Developer notifies PO that his story is ready to accept
- Developer moves story to accept

**4. From Accept to Done**<BR>
*Responsible: PO*<BR>
- PO checks all test scenarios and acceptance criteria
- Spin off stories are collected
- PO marks the story as being accepted
- Developer merges the code to master
- Developer moves story to done



### 2.5. Celebrate

Your ideas, documentation and code have become an integral part of this project. You are the Open Source hero we need.


## 3. Development process

The part below is meant as documentation for the team developing and maintaining NLX. It is public to give more insight in the development process.


### 3.1. Agile scrum

The NLX Team uses the Agile Scrum framework for product development. It encourages us to learn through experiences, self-organise while working on a problem, and reflect on our wins and losses to continuously improve.


#### 3.1.1. Sprints

The development process is structured in sprints:

* A sprint starts on Wednesdays and lasts for two weeks.
* We plan most scrum rituals on that Wednesday:
  * *Sprint review* - review sprint results and adapt backlog
  * *Retrospective* - improving the team with every sprint
  * *Refinement* - finishing touch getting user stories clear, including scrum poker
  * *Sprint planning* - starting the new sprint with a sprint goal
* The other Wednesday, half-way sprint, we spent some time refining user stories as well.

#### 3.1.2. Scrum boards

To keep track of the work, we use the [issue boards](https://docs.gitlab.com/ee/user/project/issue_board.html) of Gitlab:

* [Refinement board](https://gitlab.com/commonground/nlx/nlx/-/boards/871734?milestone_title=No+Milestone&), to prepare the next sprint:
  * Column "Open" is the backlog, filled with new user stories. The Product Owner prioritizes the backlog and selects user stories that come up next.
  * Column "Refinement" contains prioritized user stories that we are currently refining. The Product Owner takes care of a clear functional description and acceptance criteria. The team enhances on this with a proposed technical solution and complexity points (the "weight").
  * Column "Ready for Sprint" contains user stories that are refined, estimated and can be picked up in a sprint. Mostly this will happen during a sprint during Sprint planning session. When a sprint is done before the end, team members can grab new user stories from this column. Stories move from this board to the next by adding a Milestone with the name of the sprint to it.

* [Sprint board](https://gitlab.com/commonground/nlx/nlx/-/boards/871691?milestone_title=Sprint%2015&), which guides us through the sprint:
  * This board contains the sprint backlog, the commited estimation and scope of the Sprint Goal, as planned during the sprint planning.
  * How stories flow through this board is more thoroughly described in [3.2 Development flow](#3-2-development-flow).
  * Column "Open" contains user stories that are waiting to be worked on.
  * Column "Doing" shows work that someone is working on. All those stories have someone assigned.
  * Column "Review" contains user stories that can be reviewed by another team member. This includes a code review and test of functionality. We add just user stories to the board, they are linked to the involved Merge Requests.
  * Column "Accept" means the user story is presented to the Product Owner. When accepted, the changes will go to production.

There are some supporting repositories surrounding the main NLX repository. Those do not contain issues nor project boards.


### 3.2 Development flow

The development flow describes how we bring user stories from idea to production.

#### 3.2.1. Overview

Development follows this flow:

1. Add to backlog
1. Select, refine, estimate and plan
1. Code
1. Review
1. Acceptance by Product Owner
1. Merge to master branch
1. Automated deploy to acceptance environment
1. Generate version tag (manual trigger)
1. Deploy release to demo, preprod and prod (manual trigger)

Issues can be created by anyone and start at the backlog.


#### 3.2.2. Select, refine, estimate and plan

The Product Owner selects user stories from the backlog and refines them, when necessary with the team. Then the team estimates how complex it is to fulfil a story, adding the complexity points. During the sprint planning, the Product Owner and team plan stories to work towards a single sprint goal to be completed during the coming sprint. The scope of the sprint goal is negotiated by selecting a set of issues that together complete the sprint goal. Full focus and commitment goes to completing the sprint goal. During daily standups, progress is discussed.

When the sprint has started, developers select the issue that they will work from the sprint board by assigning it to themselves and by moving it from column "Open" to "Doing".


#### 3.2.3. Code

Every team member develops on a local copy of the repository. New features are added on a feature branch. During this stage, local builds and tests are made. Once ready, the feature branch is pushed to the repository hosted by Gitlab.

With the feature branch now available on Gitlab, a Merge Request is created. This is always a request to merge the new or altered code into the (default) Master branch.

Merge Requests to the Master branch trigger a CI pipeline to perform unit tests and to build all containers. If this fails, the developer will continue coding and pushing to the feature branch on Gitlab until the pipeline passes.

By starting the title of a MR with `WIP:`, one can indicate "Work in Progress". Gitlab will prevent merging work that is marked `WIP:`. This is useful to trigger a pipeline or to start a discussion.


#### 3.2.4. Review

Once ready, the developer asks for a review of the work by moving the related issue on the Sprint board from column "Doing" to "Review". It is also useful to ask for the Review by announcing it on Slack.

One or more other developers will perform a code review, commenting and discussing until everything is clear. If this results in necessary changes, the issue moves back to the column "Open" (when other work was started meanwhile) or "Doing" (when work continues) until the issue is ready for review again.

The review makes sure that all code is seen by multiple people. This prevents all sort of mistakes, makes sure knowledge is shared throughout the team and makes sure more people feel responsible about the code.

Once a reviewer is satisfied he or she will approve the Merge Request. At lease one approval is required to continue. After approval the
Product Owner accepts the story by also approving the Merge Request. The developer is responsible for merging the Merge Request to master.

Small merge requests that do not change the behaviour of the software itself (e.g. dependency updates) do not have to be accepted by the Product Owner. They can be merged after code review.

The required approval from the Product Owner can be removed by editing the Merge Request and setting the number of approvals required for PO Accept to 0.

While an issue is in 'Review', it remains assigned to the developer who is working on the issue (not the reviewer), and it is the responsibility of this assignee to make sure a timely and complete review of the proposed changes.

As we want to deliver value to the customer as soon as possible, stories shouldn't stay in Review too long. The author of the story should actively reach out to the team members to get the work reviewed.

When a branch begins with `review/` a review app is created for that branch. This app can be used to inspect UI changes of the Merge Request.


#### 3.2.5. Deploy to acceptance environment

A successful merge triggers the pipeline. After testing, it releases the build containers and automatically deploys to the acceptance environment.

After deployment to the acceptance environment, the newly deployed features are checked online by the developer. If everything still works and the new features perform as intended, the issue is moved to the 'Closed' column. If not, bug fixing is in order.


#### 3.2.6. Version tag

A new version is tagged by running the manual Semantic Release job on the master branch. First a 'dry run' job is triggered manually to check the new version number and the generated changelog.

When the output is as expected, the real release job is triggered. The job updates the [CHANGELOG.md](CHANGELOG.md), commits the changes and adds a new tag to the commit with the new version number.


#### 3.2.7. Deploy to production environment and release Docker images

A new production release is done by trigging the manual release job on the master branch. The release is deployed to the following environments:

* Demo
* Pre-production
* Production

After deployment, a manual check is done to check if everything still works as intended.

#### 3.2.8. Version skew policy

For the internal interfaces of the NLX system we have a version skew policy for MAJOR/MINOR releases of `n-2`. So for example a component of version 1.5 is able to communicate with a component of version 1.3 and a component of version 2.0 is able to communicate with the last two minor versions of the 1.x series.


### 3.3. Communication

#### 3.3.1. Gitlab
  * User story related communication discussion is mostly done in the comments below user stories
  * Reviews are done as comments below Merge Requests
  * All communication on Gitlab is written in English

#### 3.3.2. Slack
  * General communication from chit chat to important notifications is done via Slack
  * Alerts generated from operations are communicated via Slack
  * The Slack workspace is reserved for the team
  * Slack is high traffic but topics strictly separated in channels

#### 3.3.3. Whereby.com
  * Since the team does not work in one location every day, we organise our stand ups via video calls. We use https://whereby.com for this
  * Stand ups last 15 - 20 minutes. We focus on sharing what every did and what will be done that day, with the occasional exchange about impediments
  * In case the Sprint Backlog needs to be renegotiated (i.e. the scope of the sprint goal is changed), this is done during the standup
  * Whereby.com is used for one-on-one communication between team members as well

---

For more information on how to use and contribute to this project, please read the [`README`](README.md).
