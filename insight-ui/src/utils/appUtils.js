/*
    Utility functions
    v.0.0.1
*/

/**
 * Log key...val pairs in log group
 * params:
 *  title:REQUIRED
 *  ...{}
 */
const logGroup = data =>{
    let props = Object.keys(data);
    if (data.title){
        console.group(data.title);
    } else {
        console.group("logGroup");
    }

    props.map((key)=>{
        let val = data[key];
        return console.log(key,"...",val);
    });
    console.groupEnd();
}

/**
 * Prepare raw table data
 * params:
 *  colDef: array of objects (according to mui table defs)
 *  rawData: array of objects
 */
const prepTableData = ({colDef, rawData}) =>{
    //debugger
    let tableData = rawData.map((row, rid)=>{
        let rowData = {};
        //add row id
        rowData['id'] = rid
        //get col data
        for (let c in colDef){
            let col = colDef[c]
            if (row[col.src]){                 
                //format data based on type
                //extend if needed with additional types
                switch(col.type.toLowerCase()){
                    case "date":                        
                        rowData[col.id] = new Date(row[col.src]).toLocaleDateString()  
                        break;                      
                    case "time":                        
                        rowData[col.id] = new Date(row[col.src]).toLocaleTimeString()
                        break;
                    default:
                        rowData[col.id] = row[col.src]
                }                
            } else {
                rowData[col.id] = null
            }
        }
        return rowData
    })
    //return prepared table data
    return tableData;
}

const sortTableData = (array, cmp) =>{
    const stabilizedThis = array.map((el, index) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = cmp(a[0], b[0]);
        if (order !== 0) return order;
        return a[1] - b[1];
    });
    return stabilizedThis.map(el => el[0]);
}

export {
    logGroup,
    prepTableData,
    sortTableData
}