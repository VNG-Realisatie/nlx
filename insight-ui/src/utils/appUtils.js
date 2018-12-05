/**
 * Prepare raw table data
 * params:
 *  colDef: array of objects (according to mui table defs)
 *  rawData: array of objects
 */
const prepTableData = ({ colDef, rawData }) => {
    let tableData = rawData.map((row, rid) => {
        let rowData = {}

        rowData['id'] = rid

        for (let c in colDef) {
            let col = colDef[c]
            let src = col.src.split('.')
            if (row[src[0]]) {
                let val = extractValue(src, row)
                switch (col.type.toLowerCase()) {
                    case 'date':
                        rowData[col.id] = new Date(val).toLocaleDateString()
                        break
                    case 'time':
                        rowData[col.id] = new Date(val).toLocaleTimeString()
                        break
                    default:
                        rowData[col.id] = val
                }
            } else {
                rowData[col.id] = null
            }
        }
        return rowData
    })
    // return prepared table data
    return tableData
}

const extractValue = (src, row) => {
    let value = null
    if (src.length === 1) {
        if (row[src[0]]) {
            value = row[src[0]]
        }
    } else {
        value = extractValue(src.slice(1), row[src[0]])
    }
    return value
}

const sortTableData = (array, cmp) => {
    const stabilizedThis = array.map((el, index) => [el, index])
    stabilizedThis.sort((a, b) => {
        const order = cmp(a[0], b[0])
        if (order !== 0) return order
        return a[1] - b[1]
    })
    return stabilizedThis.map((el) => el[0])
}

/**
 * Extract error status (number) and description (string)
 * from error object. There are few variations how errors
 * are returned.
 * @param {Object} error
 * @returns {number} error.status
 * @returns {string} error.description
 */
const extractError = (error) => {
    let err = {
        status: null,
        description: null,
    }
    if (error.status) {
        err.status = error.status
    }
    if (error.statusText) {
        err.description = error.statusText
    }
    if (error.description) {
        err.description = error.description
    }
    if (error.data && error.data.status) {
        err.status = error.data.status
    }
    if (error.data && error.data.description) {
        err.status = error.data.description
    }
    if (error.response && error.response.status) {
        err.status = error.response.status
    }
    if (error.response && error.response.statusText) {
        err.description = error.response.statusText
    }
    if (error.response && error.response.data) {
        if (error.response.data.status) {
            err.status = error.response.data.status
        }

        if (error.response.data.statusText) {
            err.description = error.response.data.statusText
        }

        if (error.response.data.description) {
            err.description = error.response.data.description
        }
    }
    return err
}

export { prepTableData, sortTableData, extractError }
