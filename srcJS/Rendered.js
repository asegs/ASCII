import React,{Component} from "react";

class Rendered extends Component{
    constructor(props) {
        super(props);
        this.state = {
            name:this.props.match.params.name,
            body:[]
        }
    }

    async componentDidMount() {
        let data = {};
        data['name']=this.state.name;
        console.log(this.state.name)
        const response = await fetch('/viewphoto',{
            method: 'PUT',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data),
        })
        const body = await response.json();
        console.log(body);
        this.setState({body:body.split("\n")});
    }

    render() {
        let textGrid = [];
        for (let step=0;step<this.state.body.length;step++){
            textGrid.push(<p style={{color:'black',fontSize:'3px',marginTop:'0px',marginBottom:'0px',fontFamily:'courierNew, monospace',whiteSpace: 'pre'}}>{this.state.body[step]}</p>);
        }
        return(
          <div>
              {textGrid}
          </div>
        );
    }
}

export default Rendered;
