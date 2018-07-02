import React from 'react';
import {
  StyleSheet,
  Text,
  View,
  Dimensions,
  Alert,
  Button,
  Picker,
  ActionSheetIOS,
  TouchableOpacity,
  TextInput,
  Switch
} from 'react-native';
import RadioForm, {RadioButton, RadioButtonInput, RadioButtonLabel} from 'react-native-simple-radio-button';
import PRE from './PRE';
import Tag from './Tag';
const v_logo = require('./logos/v.png')
const s_logo = require('./logos/s.png')
const mp_logo = require('./logos/mp.png')
const c_logo = require('./logos/c.png')
const l_logo = require('./logos/l.png')
const m_logo = require('./logos/m.png')
const kd_logo = require('./logos/kd.png')
const sd_logo = require('./logos/sd.png')
const person_logo = require('./logos/person.png')

const WIDTH = Dimensions.get("window").width;
const HEIGHT = Dimensions.get("window").height;

const SIZE = WIDTH < HEIGHT ? WIDTH - 40 : HEIGHT - 40;

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      parties: defaultParties,
      user_choice: 0,
      start_screen: false,
      comment: "",
      is_partist: false
    };
    this.updateParty = this.updateParty.bind(this);
  }

  displayState() {
    let url = "https://galtanapi.aiman.space/results";
    let parties = {}
    Object.keys(this.state.parties).forEach(p => {
      //console.log(this.state.parties[p].x)
      parties[p] = {
        "right_left": normalize(this.state.parties[p].x),
        "gal_tan": normalize(this.state.parties[p].y)
      }
    })
    let data = {
      "political_views": parties,
      "user_choice": this.state.user_choice,
      "is_partist": this.state.is_partist,
      "comment": this.state.comment
    }
    fetch(url, {
      method: "POST",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })
    .then(status => {
      console.log(status);
      this.setState({start_screen: true});
    })
    .catch((error) =>{
      console.error(error);
    });
  }

  start() {
    this.setState({start_screen: false, parties: defaultParties, comment: "", is_partist: false})
  }

  restart() {
    this.setState({start_screen: true})
  }

  updateParty(obj, party) {
    let tempData = Object.assign({}, this.state.parties);
    tempData[party] = obj[party];
    this.setState({parties: tempData});
  }

  toggleIsPartist() {
    let is = this.state.is_partist;
    this.setState({is_partist: !is});
  }

  render() {
    return (
      <View>
        {this.state.start_screen
          ?
          <View>
            <TouchableOpacity style={styles.startButton} onPress={() => this.start()}>
              <Text style={styles.startButtonText}>START</Text>
            </TouchableOpacity>
          </View>
          :
          <View>
            

            <View style={styles.container}>
              {([0,1,2,3,4,5,6,7,8,9,10].map((x, i) => <View style={bars(i).latitude} key={i}></View>))}
              {([0,1,2,3,4,5,6,7,8,9,10].map((x, i) => <View style={bars(i).longitude} key={i}></View>))}
              
              <Tag updateParty={this.updateParty} party={"person"} logo={person_logo} />
              <Tag updateParty={this.updateParty} party={"v"} logo={v_logo} />
              <Tag updateParty={this.updateParty} party={"sd"} logo={sd_logo} />
              <Tag updateParty={this.updateParty} party={"s"} logo={s_logo} />
              <Tag updateParty={this.updateParty} party={"mp"} logo={mp_logo} />
              <Tag updateParty={this.updateParty} party={"m"} logo={m_logo} />
              <Tag updateParty={this.updateParty} party={"l"} logo={l_logo} />
              <Tag updateParty={this.updateParty} party={"kd"} logo={kd_logo} />
              <Tag updateParty={this.updateParty} party={"c"} logo={c_logo} />
              
            </View>

            <View style={{flexDirection: "row", marginTop: 30, marginLeft: 10}}>
              <RadioForm
                radio_props={radio_props}
                initial={0}
                formHorizontal={true}
                labelHorizontal={false}
                buttonColor={'rgb(150,100,200)'}
                selectedButtonColor={"rgb(150,100,200)"}
                animation={true}
                onPress={(value) => {this.setState({user_choice: value})}} />

              <View style={{flexDirection: "row"}}>
                <Text style={{margin: 10, marginLeft: 40}}>Är du aktiv i detta parti?</Text>
                <Switch
                  onValueChange={() => this.toggleIsPartist()}
                  value={this.state.is_partist}/>
              </View>
            </View>

            <View style={{borderWidth: 1, borderColor: "black", marginLeft: 20, marginRight: 20, borderRadius: 10, marginTop: 10}}>
              <TextInput
                style={{fontSize: 24, height: 40}}
                placeholder="Kommentar"
                onChangeText={(text) => this.setState({comment: text})} />
            </View>

            <View style={{flexDirection: "row", justifyContent: "center", marginTop: 10}}>
              <TouchableOpacity style={styles.doneButton} onPress={() => this.restart()}>
                <Text style={styles.doneButtonText}>OMSTART</Text>
              </TouchableOpacity>
              <TouchableOpacity style={styles.doneButton} onPress={() => this.displayState()}>
                <Text style={styles.doneButtonText}>KLAR</Text>
              </TouchableOpacity>
            </View>

          </View>
        }
      </View>
    );
  }
}

var radio_props = [
  {label: "Vet ej", value: 0 },
  {label: 'C', value: 1 },
  {label: 'L', value: 2 },
  {label: 'KD', value: 3 },
  {label: 'M', value: 4 },
  {label: 'MP', value: 5 },
  {label: 'S', value: 6 },
  {label: 'SD', value: 7 },
  {label: 'V', value: 8 },
  {label: "Annat", value: 9 }
];

const styles = StyleSheet.create({
  container: {
    backgroundColor: '#ffa',
    alignItems: 'center',
    justifyContent: 'center',
    width: SIZE,
    height: SIZE,
    marginTop: 30,
    marginLeft: 20
  },
  image: {
    width: 'auto',
    height:''
  },
  startButton: {
    backgroundColor: "rgb(150,100,200)",
    height: 100,
    width: 200,
    marginLeft: "auto",
    marginRight: "auto",
    marginTop: HEIGHT/2-100,
    borderRadius: 20,
    shadowColor: "black",
    shadowOffset: {width: 0, height: 5},
    shadowOpacity: 0.7,
    shadowRadius: 4,
    justifyContent: "center"
  },
  startButtonText: {
    textAlign: "center",
    fontSize: 36,
    color: "ivory",
    fontWeight: "200"
  },
  doneButton: {
    backgroundColor: "rgb(150,100,200)",
    height: 60,
    width: 180,
    margin: 20,
    borderRadius: 20,
    shadowColor: "black",
    shadowOffset: {width: 0, height: 5},
    shadowOpacity: 0.7,
    shadowRadius: 4,
    justifyContent: "center"
  },
  doneButtonText: {
    textAlign: "center",
    fontSize: 36,
    color: "ivory",
    fontWeight: "200"
  }
});

const bars = index => {
  return StyleSheet.create({
    latitude: {
      width: SIZE,
      height: 2,
      backgroundColor: index === 0 || index === 5 || index === 10 ? "#222" : "#bbb",
      //transform: [{ translateY: (-SIZE/2) + ((SIZE*index)/10 - index) + 5 }]
      position: "absolute",
      top: (SIZE/10)*index
    },
    longitude: {
      width: 2,
      height: SIZE,
      backgroundColor: index === 0 || index === 5 || index === 10 ? "#222" : "#bbb",
      position: "absolute",
      left: (SIZE/10)*index
    }
  })
}

const normalize = (n) => {
  return n !== 0 ? n._value/(SIZE/2) : 0;
}

const defaultParties = {
  v: {x: 0, y: 0},
  s: {x: 0, y: 0},
  mp: {x: 0, y: 0},
  c: {x: 0, y: 0},
  l: {x: 0, y: 0},
  m: {x: 0, y: 0},
  kd: {x: 0, y: 0},
  sd: {x: 0, y: 0}
}