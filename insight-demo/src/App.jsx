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
    before: { opacity: 0, x: '100%' },
    enter: { opacity: 1, x: 0,
        transition: {
            default: { type: 'spring', stiffness: 35, damp: 200, duration: 150 }
        }
    },
    exit: { opacity: 0, x: '-100%' },
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
    padding-bottom: 8rem;
`

const StyledContent = styled.div`
    width: 100%;
    max-width: 520px;
    margin: 0 auto;
`

const buttonStyle = {
    position: 'absolute',
    top: 0,
    bottom: 0,
    margin: 'auto 0',
    zIndex: 1
}

const backButtonStyle = {
    left: '4rem'
}

const nextButtonStyle = {
    right: '4rem'
}

class App extends Component {
	render() {
		return (
			<ThemeProvider theme={theme}>
				<Route render={({ location }) => {
                    let prevLink = null
                    let nextLink = null
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
                                    {prevLink &&
                                        <Button variant="tertiary" as={Link} to={prevLink} style={{...buttonStyle,...backButtonStyle}}>Back</Button>
                                    }
                                    {nextLink &&
                                        <Button as={Link} to={nextLink} style={{...buttonStyle,...nextButtonStyle}}>Next</Button>
                                    }
                                    <StyledContent>
                                        <PoseGroup>
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
