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
We follow the [Angular Commit Message conventions](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines). This convention requires you to pas a subject and scope in the commit message. The scope is based on the applications in the repository. If you are not sure which scope to use please leave the scope empty.

The available scopes are:

- auth-service
- ca-certportal
- ca-cfssl-unsafe
- common
- design
- directory-db
- directory-monitor
- directory-inspection-api
- directory-registration-api
- directory-ui
- docs
- helm
- insight-api
- insight-demo
- insight-ui
- inway
- outway
- txlog-db

#### 2.2.4. Do NOT close issues with commit messages
Make sure _not_ to use the commit message to [automatically close issues](https://docs.gitlab.com/ee/user/project/issues/automatic_issue_closing.html), since we do _not_ want issues to be closed immediately after merging to the master branch.

### 2.3. Merge Request

#### 2.3.1. Always refer to an issue
Before starting a Merge Request, make sure there is a User Storiy describing what you want to achieve with the MR. [Create a story by submitting a new issue](https://gitlab.com/commonground/nlx/nlx/issues) if there is none. New issues come with a User Story template. This template helps you think from the user perspective: 'who wants this new feature and why?'

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

With MR's we make User Stories become reality. A User Story is DONE when:
* Code has been written.
* Documentation has been added or updated where necessary.
* All changes have been reviewed and approved by another developer from [team NLX](https://gitlab.com/commonground/nlx/nlx/-/project_members).
* Deliverables (functionality + documentation) have been demoed to the development team and any resulting feedback has been processed.
* Deliverables (functionality + documentation) have been demoed to the Product Owner and any resulting feedback has been processed.
* Any spin-off user stories have been clearly identified and brought to the attention of the Product Owner.
* Sprint demo has been prepared (test data, scenarios, etc.).
* Product owner has accepted the user story.
* All changes are deployed to the production environments.


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

Development follows a flow:

1. Add to backlog
2. Select, refine, estimate and plan
3. Code
4. Review
5. Merge to master branch
6. Deploy to test environment
7. Create a versioned release
8. Deploy to acceptance environment
9. Acceptance by Product Owner
10. Deploy to production environment

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

Once a reviewer is satisfied he or she will approve the Merge Request. At lease one approval is required to continue.

While an issue is in 'Review', it remains assigned to the developer who is working on the issue (not the reviewer), and it is the responsibility of this assignee to make sure a timely and complete review of the proposed changes. Stories shouldn't stay in Review too long.


#### 3.2.5. Merge

*Note: this part is identified as sub optimal. In the current setup, the best we can do is "Move fast and break things" (because code is merged before it is accepted). If we want to improve this we need review apps or release channels.*

With the approval of the automatic tests from the pipeline and the human code review, the Merge Request is now ready to be merged. Gitlab will refuse to merge without those "green flags".

Sometimes a merge cannot be done automatically because it contains commits that touch lines of code that were altered by another merge, resulting in a "merge conflict" that a developer can resolve manually.


#### 3.2.6. Deploy to test environment

A successful merge triggers another pipeline, which again runs unit tests. Then it releases the build containers and deploys them to the test environment.

After deployment to test environment, the newly deployed features are checked online. If everything still works and the new features perform as intended, a new issue is selected to work on. If not, bug fixing is in order.


#### 3.2.7. Create a versioned release

If the test environment looks OK, the deployment should move to acc (acceptance) environment where the new or changed functionality can be accepted.

For this, a versioned release is needed. The developer can manually trigger the Semantic Release tool. This tool looks at all commit messages on the Master branch since the latest version, and creates a new one. All commit messages are parsed and a new version number is generated. Depending on the commit messages, the major, minor or patch number is increased.


#### 3.2.8. Deploy to acceptance environment

With the versioned release, the same pipeline as for test deployment is fired again, this time to deploy to acc environment.

After deployment to acc, the features are checked online. When OK, the issue on the board is moved from column "Review" to "Acceptance".


#### 3.2.9. Acceptance by Product Owner

All issues in the column "Acceptance" are reviewed by the Product Owner. If the acceptance criteria are met and definition of done is followed, the issue is accepted and moved to the column "closed".


#### 3.2.10. Deploy to production environment

Once accepted, the release can be deployed to the following environments:

* Demo
* Preprod
* Prod

These three environments should be at the same version at all times.

Deployment is triggered manually by the Product Owner. After deployment, a manual check is done to check if everything still works as intended.

If so, the issue is moved to the column "Closed".


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

#### 3.3.3. Appear.in
  * Since the team does not work in one location every day, we organise our stand ups via video calls. We use https://appear.in for this
  * Stand ups last 15 - 20 minutes. We focus on sharing what every did and what will be done that day, with the occasional exchange about impediments
  * In case the Sprint Backlog needs to be renegotiated (i.e. the scope of the sprint goal is changed), this is done during the standup
  * Appear.in is used for one-on-one communication between team members as well

---

For more information on how to use and contribute to this project, please read the [`README`](README.md).
