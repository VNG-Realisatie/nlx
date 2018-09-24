
export const config = {
  appTitle:"This is app title",
  logo:{
    src:"" 
  },
  api:{
    baseUri:"http://localhost:3000/api/",
    listServices:()=>{
      let url = "/directory/list-services";
      return url;
    },
  }
}

export default config;