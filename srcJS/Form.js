import React,{Component} from 'react'
import Slider from '@material-ui/core/Slider';
import calculate from "./Dimensions";
import PhotoElement from "./PhotoElement";

class Form extends Component{
    constructor(props) {
        super(props);
        this.state = {
            addressMode:false,
            latitude:43.0844955,
            longitude:-77.6749311,
            maxShift:calculate(43.0844955,15),
            zoom:15,
            address:"",
            inverse:false,
            photo:undefined,
            shift:(0.33*calculate(43.0844955,15)).toFixed(3),
            text:[],
            photoNames:[]
        }
        this.submitCoords = this.submitCoords.bind(this);
        this.submitAddress = this.submitAddress.bind(this);
        this.uploadPhoto = this.uploadPhoto.bind(this);
    }

    changeLat =(e)=>{
        this.setState({latitude:e.target.value},()=>{
            this.setState({maxShift:calculate(this.state.latitude,this.state.zoom)})
        });
    }

    changeLon = (e)=>{
        this.setState({longitude:e.target.value});
    }

    changeZoom = (e)=>{
        this.setState({zoom:e.target.value},()=>{
            this.setState({maxShift:calculate(this.state.latitude,this.state.zoom)})
        });
    }

    changeAddress = (e)=>{
        this.setState({address:e.target.value});
    }

    toggleMode = ()=>{
        this.setState({addressMode:!this.state.addressMode});
    }

    toggleInvert = ()=>{
        this.setState({inverse:!this.state.inverse});
    }

    updateFile =(e)=>{
        this.setState({photo:e.target.files[0]});
    }

    updateMagnitude=(event,value)=>{
        this.setState({shift:value});
    }

    zoomOut=()=>{
        if (this.state.zoom<=1){
            return;
        }
        this.setState({zoom:parseInt(this.state.zoom)-1},()=>{
            this.submitCoords();
        })
    }

    zoomIn=()=>{
        if (this.state.zoom>=22){
            return;
        }
        this.setState({zoom:parseInt(this.state.zoom)+1},()=>{
            this.submitCoords();
        })
    }

    async submitCoords(){
        let numericalLat = parseFloat(this.state.latitude);
        let numericalLong = parseFloat(this.state.longitude);
        let numericalZoom = parseInt(this.state.zoom);
        let data = {};
        data['latitude'] = numericalLat;
        data['longitude'] = numericalLong;
        data['zoom'] = numericalZoom;
        data['inverse'] = this.state.inverse;
        const response = await fetch('/changelocation',{
            method: 'PUT',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data),
        })
        const body = await response.json();
        console.log(body);
        if (body['body']){
            this.setState({text:body['body'].split("\n"),maxShift:calculate(this.state.latitude,this.state.zoom),photoNames:body['photos']},()=>{
                this.setState({shift:(0.33*this.state.maxShift).toFixed(3)});
            });
        }
        else {
            this.setState({text: body.split("\n"), maxShift: calculate(this.state.latitude, this.state.zoom)},()=>{
                this.setState({shift:(0.33*this.state.maxShift).toFixed(3)});
            });
        }
    }

    async submitAddress(){
        let numericalZoom = parseInt(this.state.zoom);
        let data = {};
        data['address'] = this.state.address;
        data['zoom'] = numericalZoom;
        data['inverse'] = this.state.inverse;
        const response = await fetch('/changeaddress',{
            method: 'PUT',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data),
        })
        const body = await response.json();
        console.log(body)
        if (body['body']){
            this.setState({text:body['body'].split("\n"),latitude:body['latitude'],longitude:body['longitude'],addressMode:false,photoNames:body['photos']},()=>{
                this.setState({maxShift:calculate(this.state.latitude,this.state.zoom)},()=>{
                    this.setState({shift:(0.33*this.state.maxShift).toFixed(3)});
                });
                document.getElementById("mode").checked = false;
            });
        }
        else {
            this.setState({text: body.split("\n"), maxShift: calculate(this.state.latitude, this.state.zoom)},()=>{
                this.setState({shift:(0.33*this.state.maxShift).toFixed(3)});
            });
        }
    }

    shiftDown=()=>{
        this.setState({latitude:parseFloat(this.state.latitude)-this.state.shift},()=>{
            this.submitCoords();
        });
    }

    shiftUp=()=>{
        this.setState({latitude:parseFloat(this.state.latitude)+this.state.shift},()=>{
            this.submitCoords();
        });
    }

    shiftLeft=()=>{
        this.setState({longitude:this.state.longitude-this.state.shift},()=>{
            this.submitCoords();
        });
    }

    shiftRight=()=>{
        this.setState({longitude:this.state.longitude+this.state.shift},()=>{
            this.submitCoords();
        });
    }

    async uploadPhoto(){
        let formData = new FormData;
        formData.append("latitude",this.state.latitude);
        formData.append("longitude",this.state.longitude);
        formData.append("file",this.state.photo);
        formData.append("extension",this.state.photo.name.substring(this.state.photo.name.indexOf(".")));
        console.log(formData.get("extension"));
        console.log(formData.get("file"))
        await fetch('/uploadphoto',{
            method:'POST',
            headers: {
                'Accept': 'application/json',
            },
            body: formData,
        })
        this.submitCoords();
    }

    handler = () =>{
        if (this.state.addressMode){
            this.submitAddress();
        }
        else{
            this.submitCoords();
        }
    }

    jumpTo = (lat, lng)=>{
        this.setState({latitude:lat,longitude:lng,zoom:20},()=>{
            this.submitCoords();
        })
    }

    render() {
        let textGrid = [];
        for (let step=0;step<this.state.text.length;step++){
            textGrid.push(<p style={{color:'black',fontSize:'3px',marginTop:'0px',marginBottom:'0px',fontFamily:'courierNew, monospace',whiteSpace: 'pre'}}>{this.state.text[step]}</p>);
        }

        let photos = [];
        for (let step=0;step<this.state.photoNames.length;step++){
            photos.push(<PhotoElement name={this.state.photoNames[step]['name']} longitude={this.state.photoNames[step]['longitude']} latitude={this.state.photoNames[step]['latitude']} jumpTo={this.jumpTo}/>);
        }

        let submit = "";
        if (this.state.photo!==undefined){
            submit = <button onClick={this.uploadPhoto} >Upload photo!</button>

        }
        return(
          <div>
              Address mode: <input id={"mode"} type={"checkbox"} onChange={this.toggleMode}/><br/>
              Invert color: <input type={"checkbox"} onChange={this.toggleInvert}/><br/>
              Latitude: <input onChange={this.changeLat} value={this.state.latitude} type={this.state.addressMode ? "hidden" : "text"}/><br/>
              Longitude: <input onChange={this.changeLon} value={this.state.longitude} type={this.state.addressMode ? "hidden" : "text"}/><br/>
              Address: <input onChange={this.changeAddress} type={this.state.addressMode ? "text" : "hidden"}/><br/>
              Zoom: <input onChange={this.changeZoom} value={this.state.zoom} type={"number"}/><br/>

              <button style={{backgroundColor: '#85bcea',padding: '10px 32px',cursor: 'pointer',borderRadius:"10px"}} onClick={this.handler}>Render</button><br/>
              <div style={{width:'20%',margin:'auto'}}>
        <Slider
            defaultValue={20}
            step={0.0001}
            valueLabelDisplay="auto"
            onChangeCommitted={this.updateMagnitude}
            min={0}
            max={this.state.maxShift}
        />
        <p>Shift: &#177;{this.state.shift}&#176;</p>
              </div>
              <button onClick={this.shiftLeft} style={{float:'left',marginLeft:'15%'}}>{"<="}</button>
              <button onClick={this.zoomOut} style={{float:'center-left',marginLeft:'30%'}}>{"(-)"}</button>
              <button onClick={this.shiftUp} style={{float:'center'}}>{"^"}</button>
              <button onClick={this.zoomIn} style={{float:'center-right',marginRight:'30%'}}>{"(+)"}</button>
              <button onClick={this.shiftRight} style={{float:'right',marginRight:'15%'}}>{"=>"}</button>
              {textGrid}<br/>
              <button onClick={this.shiftDown} style={{float:'center'}}>{"v"}</button><br/>
              Photo: <input type={"file"} onChange={this.updateFile}/>
              {submit}
              {photos}
          </div>
        );
    }
}

export default Form;

