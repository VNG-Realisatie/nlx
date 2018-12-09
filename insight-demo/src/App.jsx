import React, { Fragment, Component } from "react"
import { Route, Link, Switch } from 'react-router-dom';
import posed, { PoseGroup } from 'react-pose';
import { media } from './theme/helpers'

import styled, { ThemeProvider } from 'styled-components'
import theme from './theme/themeConstants'
import GlobalStyle from './globalStyle'
import { Flex, Box } from '@rebass/grid'
import { Container } from './components/Grid/Grid'

import Button from "./components/Button/Button";
import Stepper from "./components/Stepper/Stepper";
import logo from './images/logo.svg'

// Pages
import Intro from './pages/Intro'
import StepOne from './pages/Step1'
import StepTwo from './pages/Step2'
import StepThree from './pages/Step3'
import StepFour from './pages/Step4'

const Header = styled.div`
    ${media.xsDown`
        padding-top: 32px;
        padding-bottom: 40px;
    `}

    ${media.xsUp`
        padding-top: 80px;
        padding-bottom: 32px;
    `}
`

const Logo = styled.div`
    ${media.xsDown`
        margin-bottom: 24px;
    `}

    ${media.xsUp`
        position: absolute;
        left: 0;
        right: 0;
        bottom: 1.75rem;
        margin: 0 auto;
    `}

    display: flex;
    justify-content: center;

    img {
        height: 16px;
    }
`

const RouteContainer = posed.div({
    before: { opacity: 0, x: ({ direction }) => direction === 'next' ? '100%' : '-100%',
    },
    enter: { opacity: 1, x: 0,
        transition: {
            default: {
                type: 'spring', stiffness: 30, mass: .5
            },
            opacity: {
                ease: 'easeInOut', duration: 750
            }
        }
    },
    exit: { opacity: 0, x: ({ direction }) => direction === 'next' ? '-100%' : '100%',
        transition: {
            default: {
                type: 'spring', stiffness: 30, mass: .5
            },
            opacity: {
                ease: 'easeInOut', duration: 300
            }
        }
    },
});

const StyledContainer = styled(Container)`
    position: relative;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    overflow: hidden;
`

const StyledPage = styled.div`
    flex-grow: 1;

    ${media.xsDown`
        padding-bottom: 6.5rem;
    `}

    ${media.xsUp`
        padding-bottom: 8rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
    `}
`

const StyledLink = styled(Link)`
    text-decoration: none;
`

const StyledContent = styled.div`
    flex-shrink: 1;
    max-width: 584px;

    ${media.xsUp`
        padding: 0 2rem;
    `}
`

const LgButton = styled(Button)`
    z-index: 1;

    ${media.xsDown`
        display: none;
    `}

    ${media.xsUp`
        display: inline-flex;
    `}
`

const SmButton = styled(Button)`
    z-index: 1;
`

const Toolbar = styled(Flex)`
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0%;
    padding: 12px 0 14px;
    background-color: white;
    box-shadow: rgba(0,0,0,0.03) 0px -1px 6px 0px, rgba(0,0,0,0.03) 0px -20px 20px -20px;

    ${media.xsUp`
        display: none;
    `}
`

class App extends Component {
    constructor(props) {
        super(props)

        this.state = {
            direction: 'next'
        }
    }

    setDirection = (direction) => {
        this.setState({direction})
    }

	render() {
		return (
			<ThemeProvider theme={theme}>
				<Route render={({ location }) => {
                    let prevLink
                    let nextLink
                    switch (location.pathname) {
                        case '/':
                            prevLink = null
                            nextLink = "/stepone"
                            break
                        case '/stepone':
                            prevLink = "/"
                            nextLink = "/steptwo"
                            break
                        case '/steptwo':
                            prevLink = "/stepone"
                            nextLink = "/stepthree"
                            break
                        case '/stepthree':
                            prevLink = "/steptwo"
                            nextLink = "/stepfour"
                            break
                        case '/stepfour':
                            prevLink = "/stepthree"
                            nextLink = null
                            break
                        default:
                            break
                    }

                    return (
                        <Fragment>
                            <GlobalStyle />
                            <StyledContainer>
                                <Header>
                                    <Container>
                                        <Flex alignItems="center" flexDirection="column">
                                            <Logo>
                                                <a href="https://nlx.io/" target="_blank" rel="noopener noreferrer">
                                                    <img src={logo} alt="logo" />
                                                </a>
                                            </Logo>
                                            <Stepper pathname={location.pathname} />
                                        </Flex>
                                    </Container>
                                </Header>
                                <StyledPage>
                                    {prevLink ?
                                        <StyledLink to={prevLink} onClick={() => this.setDirection('back')}>
                                            <LgButton variant="tertiary">Back</LgButton>
                                        </StyledLink>
                                        :
                                        <LgButton variant="tertiary" disabled >Back</LgButton>
                                    }
                                    <StyledContent>
                                        <PoseGroup preEnterPose="before" direction={this.state.direction}>
                                            <RouteContainer key={location.pathname}>
                                                <Switch location={location}>
                                                    <Route exact path="/" component={Intro} key="intro" />
                                                    <Route path="/stepone" component={StepOne} key="stepone" />
                                                    <Route path="/steptwo" component={StepTwo} key="steptwo" />
                                                    <Route path="/stepthree" component={StepThree} key="stepthree" />
                                                    <Route path="/stepfour" component={StepFour} key="stepfour" />
                                                </Switch>
                                            </RouteContainer>
                                        </PoseGroup>
                                    </StyledContent>
                                    {nextLink ?
                                        <StyledLink to={nextLink} onClick={() => this.setDirection('next')}>
                                            <LgButton>Next</LgButton>
                                        </StyledLink>
                                        :
                                        <LgButton disabled >Next</LgButton>
                                    }
                                    <Toolbar justifyContent="center" mt={6}>
                                        <Box mr={3}>
                                            {prevLink ?
                                                <StyledLink to={prevLink} onClick={() => this.setDirection('back')}>
                                                    <SmButton variant="tertiary">Back</SmButton>
                                                </StyledLink>
                                                :
                                                <SmButton variant="tertiary" disabled >Back</SmButton>
                                            }
                                        </Box>
                                        {nextLink ?
                                            <StyledLink to={nextLink} onClick={() => this.setDirection('next')}>
                                                <SmButton>Next</SmButton>
                                            </StyledLink>
                                            :
                                            <SmButton disabled >Next</SmButton>
                                        }
                                    </Toolbar>
                                </StyledPage>
                            </StyledContainer>
                        </Fragment>
                    )
                }} />
			</ThemeProvider>
		);
	}
}

export default App;
