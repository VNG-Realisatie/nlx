import React from 'react'
import {Container} from 'src/Grid/Grid'
import Card from 'src/Card/Card'
import Table, { TableHead, TableBody, TableRow, TableCell } from 'src/Table/Table'

let id = 0;
function createData(name, calories, fat, carbs, protein) {
    id += 1;
    return { id, name, calories, fat, carbs, protein };
}

const rows = [
    createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
    createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
    createData('Eclair', 262, 16.0, 24, 6.0),
    createData('Cupcake', 305, 3.7, 67, 4.3),
    createData('Gingerbread', 356, 16.0, 49, 3.9),
];

export const tableStory = (
    <Container mt={4}>
        <Card>
            <div style={{overflowX: 'auto'}}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell as="th">Dessert (100g serving)</TableCell>
                            <TableCell as="th" align="right">Calories</TableCell>
                            <TableCell as="th" align="right">Fat (g)</TableCell>
                            <TableCell as="th" align="right">Carbs (g)</TableCell>
                            <TableCell as="th" align="right">Protein (g)</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                    {rows.map(row => {
                        return (
                            <TableRow key={row.id}>
                                <TableCell as="th">{row.name}</TableCell>
                                <TableCell align="right">{row.calories}</TableCell>
                                <TableCell align="right">{row.fat}</TableCell>
                                <TableCell align="right">{row.carbs}</TableCell>
                                <TableCell align="right">{row.protein}</TableCell>
                            </TableRow>
                        )
                    })}
                    </TableBody>
                </Table>
            </div>
        </Card>
    </Container>
)
