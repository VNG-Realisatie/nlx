# Contributing

We – [VNG realisatie](https://www.vngrealisatie.nl/) and the maintainers of this project – know we can only build NLX together with you. Thus we appreciate your input, enjoy feedback and welcome improvements to this project and are very open to collaboration.

We love issues and merge requests from everyone.

## Problems, suggestions and questions in Issues

You don't need to change any of our code or documentation to be a contributor. Please help development by reporting problems, suggesting changes and asking questions. To do this, you can [create an issue using GitLab](https://docs.gitlab.com/ee/user/project/issues/create_new_issue.html) for this project in the [GitLab Issues for NLX](https://gitlab.com/commonground/nlx/issues).

## Documentation and code in Merge Requests

If you want to add to the documentation or code of one of our projects you should push a branch and make a Merge Request. If you have never used GitLab before, get up to speed by reading about the [GitLab Flow](https://about.gitlab.com/2014/09/29/gitlab-flow/).

### 1. Make your changes

#### 1.1. Use OneFlow
This project uses the **OneFlow branching model** and workflow. When you've forked this repository, please make sure to create a feature branch following the OneFlow model. Read this [short blogpost](http://endoflineblog.com/oneflow-a-git-branching-model-and-workflow) when you're not yet familiar with OneFlow.

#### 1.2. Add docs and tests
If you are adding code, make sure you've added and updated the relevant documentation and tests before you submit your Merge Request. Make sure to write unit tests that show the behaviour of the newly added or changed code.

### 2. Commit messages

#### 2.1. Explain your contrributions
Add your changes in commits [with a message that explains them](https://robots.thoughtbot.com/5-useful-tips-for-a-better-commit-message). Document choices or decisions you make in the commit message, this will enable everyone to be informed of your choices in the future.

#### 2.2. Semantic Release
This project uses [semantic-release](https://semantic-release.gitbook.io/semantic-release/). When merging a MR to master, this will automatically generate our [CHANGELOG](./CHANGELOG.md) based on the commit messages and a version tag will be added.

#### 2.3. Use Angular Commit Message Convention
We follow the [Angular Commit Message Convention](https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines). This convention requires you to pas a subject and scope in the commit message. The scope is based on the applications in the repository. If you are not sure which scope to use please leave the scope empty.

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
- txlog-api
- txlog-db
- txlog-ui

#### 2.4. Do NOT close issues with commit messages
Make sure _not_ to use the commit message to [automatically close issues](https://docs.gitlab.com/ee/user/project/issues/automatic_issue_closing.html), since we do _not_ want issues to be closed immediately after merging to the master branch.

### 3. Merge Request

#### 3.1. Always refer to an issue
Before starting a Merge Request, make sure there are one or more User Stories describing what you want to achieve with the MR. [Create user stories by submitting a new issue](https://gitlab.com/commonground/nlx/issues) if there are none. New issues come with a User Story template. This template helps you think from the user perspective: 'who wants this new feature and why?'

#### 3.2. Describe the MR

When submitting the Merge Request, please accompany it with a short description of the problem you are trying to address and the issue numbers that this Merge Request fixes/addresses.

#### 3.1. Combine frontend and backend work in one MR
When working on a feature which requires specific knowledge of multiple disciplines (eg. both frontend & backend), make sure to complete your MR before asking for a review.
By doing so, the reviewer can consider the complete solution and give more insightful feedback.

### 4. Improve

#### 4.1. Reviews
All contributions have to be reviewed by someone. It could be that your contribution can be merged immediately by a maintainer. However, usually, a new Merge Request needs some improvements before it can be merged. Other contributors (or our automatic testing jobs) might have feedback. If this is the case the reviewing maintainer will help you improve your documentation and code.

#### 4.2. Definition of Done

With MR's we make User Stories become reality. A User Story is DONE when:
- Code has been written
- Documentation has been added or updated where necessary
- All changes have been reviewed and approved by another developer from [team NLX](https://gitlab.com/commonground/nlx/-/project_members)
- Deliverables (functionality + documentation) have been demoed to the development team and any resulting feedback has been processed
- Deliverables (functionality + documentation) have been demoed to the Product Owner and any resulting feedback has been processed
- Known bugs/issues have been resolved
- Any spin-off user stories have been clearly identified and brought to the attention of the Product Owner
- Sprint demo has been prepared (test data, scenarios, etc.)
- Product owner has accepted the user story
- All changes are deployed to the production environments


### 5. Celebrate

Your ideas, documentation and code have become an integral part of this project. You are the Open Source hero we need.

---

For more information on how to use and contribute to this project, please read the [`README`](README.md).
