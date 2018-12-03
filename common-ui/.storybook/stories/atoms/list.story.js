import React from 'react'
import List, {ListItem} from 'src/List/List'

export const listStory = (
    <List style={{width: 140}}>
        {['<ListItem />', '<ListItem />', '<ListItem />', '<ListItem />', '<ListItem />'].map((text, index) => (
            <ListItem key={index} as="a">
                {text}
            </ListItem>
        ))}
    </List>
)
