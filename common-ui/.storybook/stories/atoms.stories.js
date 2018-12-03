import { Fragment } from 'react'
import { storiesOf } from '@storybook/react'
import { withInfo } from '@storybook/addon-info'

import { Flex, Box } from '@rebass/grid'

import { buttonStory } from './atoms/button.story.js'
import { iconbuttonStory } from './atoms/iconbutton.story.js'
import { cardStory } from './atoms/card.story.js'
import * as inputStory from './atoms/input.story.js'
import { sectionStory } from './atoms/section.story.js'
import { paragraphStory } from './atoms/paragraph.story.js'
import { titleStory } from './atoms/title.story.js'
import { smallStory } from './atoms/small.story.js'
import { checkboxStory } from './atoms/checkbox.story.js'
import { switchStory } from './atoms/switch.story.js'
import { linkStory } from './atoms/link.story.js'
import { listStory } from './atoms/list.story.js'
import { navigationStory } from './atoms/navigation.story.js'
import { drawerStory } from './atoms/drawer.story.js'
import { errorStory } from './atoms/error.story.js'

storiesOf('Atoms', module)
    .addDecorator(withInfo)
    .addParameters({
        info: {
            propTablesExclude: [Fragment, Flex, Box, Fragment],
            text: ''
        }
    })
    .add(
        'Button',
        () => buttonStory,
        { info: { text: 'Description or documentation about my component, supports markdown.' } }
    )
    .add('Card', () => cardStory)
    .add('Checkbox', () => checkboxStory)
    .add('Drawer', () => drawerStory)
    .add('Error', () => errorStory)
    .add('IconButton', () => iconbuttonStory)
    .add('Input', () => inputStory.input, { info: { text: inputStory.info } })
    .add('Link', () => linkStory)
    .add('List', () => listStory)
    .add('Navigation', () => navigationStory)
    .add('Paragraph', () => paragraphStory)
    .add('Section', () => sectionStory)
    .add('Small', () => smallStory)
    .add('Switch', () => switchStory)
    .add('Title', () => titleStory)
