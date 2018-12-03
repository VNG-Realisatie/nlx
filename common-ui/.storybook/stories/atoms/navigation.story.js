import React from 'react'
import Navigation, {NavigationItem} from 'src/Navigation/Navigation'

export const navigationStory = (
    <Navigation>
        {['<NavigationItem />', '<NavigationItem />', '<NavigationItem />', '<NavigationItem />'].map((text, index) => (
            <NavigationItem key={index} as="a">
                {text}
            </NavigationItem>
        ))}
    </Navigation>
)
