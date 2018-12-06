import React, { Fragment, Component } from "react"
import { Route, Link, Switch } from 'react-router-dom';
import posed, { PoseGroup } from 'react-pose';

import styled, { ThemeProvider } from 'styled-components'
import theme from './theme'
import GlobalStyle from './globalStyle'
import { Flex, Box } from '@rebass/grid'
import { Container } from './components/Grid/Grid'

import Button from "./components/Button/Button";
import Stepper from "./components/Stepper/Stepper";

// Pages
import Intro from './pages/Intro'
import StepOne from './pages/Step1'
import StepTwo from './pages/Step2'
import StepThree from './pages/Step3'
import StepFour from './pages/Step4'

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
`

const StyledPage = styled.div`
    flex-grow: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-bottom: 8rem;
`

const StyledContent = styled.div`
    flex-shrink: 1;
    max-width: 584px;
    padding: 0 2rem;
`

const buttonStyle = {
    zIndex: 1
}

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
                                <Box pt={6} pb={4}>
                                    <Container>
                                        <Flex justifyContent="center">
                                            <Stepper pathname={location.pathname} />
                                        </Flex>
                                    </Container>
                                </Box>
                                <StyledPage>
                                    {prevLink ?
                                        <Button variant="tertiary" as={Link} to={prevLink} style={buttonStyle} onClick={() => this.setDirection('back')}>Back</Button>
                                        :
                                        <Button variant="tertiary" disabled style={buttonStyle}>Back</Button>
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
                                        <Button as={Link} to={nextLink} style={buttonStyle} onClick={() => this.setDirection('next')}>Next</Button>
                                        :
                                        <Button disabled style={buttonStyle}>Next</Button>
                                    }
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
