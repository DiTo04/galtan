import React from 'react';
import { StyleSheet, Text, View, Dimensions, Alert, Button } from 'react-native';
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

const WIDTH = Dimensions.get("window").width;
const HEIGHT = Dimensions.get("window").height;

const SIZE = WIDTH < HEIGHT ? WIDTH - 10 : HEIGHT - 10;

export default class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      parties: {
        v: {x: 0, y: 0},
        s: {x: 0, y: 0},
        mp: {x: 0, y: 0},
        c: {x: 0, y: 0},
        l: {x: 0, y: 0},
        m: {x: 0, y: 0},
        kd: {x: 0, y: 0},
        sd: {x: 0, y: 0}
      },
      user_choice: "david_tennander"
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
    console.log(parties);
    let data = {
      "political_views": parties,
      "user_choice": this.state.user_choice
    }
    fetch(url, {
      method: "POST",
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })
    .catch((error) =>{
      console.error(error);
    });
  }

  updateParty(obj, party) {
    console.log(typeof parseFloat(obj[party].x))
    let tempData = Object.assign({}, this.state.parties);
    tempData[party] = obj[party];
    this.setState({parties: tempData});
  }

  render() {
    return (
      <View>
        <View style={styles.container}>
          {([0,1,2,3,4,5,6,7,8,9,10].map((x, i) => <View style={bars(i).latitude} key={i}></View>))}
          {([0,1,2,3,4,5,6,7,8,9,10].map((x, i) => <View style={bars(i).longitude} key={i}></View>))}
          <Tag updateParty={this.updateParty} party={"v"} logo={v_logo} />
          <Tag updateParty={this.updateParty} party={"s"} logo={s_logo} />
          <Tag updateParty={this.updateParty} party={"mp"} logo={mp_logo} />
          <Tag updateParty={this.updateParty} party={"c"} logo={c_logo} />
          <Tag updateParty={this.updateParty} party={"l"} logo={l_logo} />
          <Tag updateParty={this.updateParty} party={"m"} logo={m_logo} />
          <Tag updateParty={this.updateParty} party={"kd"} logo={kd_logo} />
          <Tag updateParty={this.updateParty} party={"sd"} logo={sd_logo} />
        </View>
        <Button onPress={() => this.displayState()} title="lol" />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: '#ffa',
    alignItems: 'center',
    justifyContent: 'center',
    width: SIZE,
    height: SIZE,
    marginTop: 100,
    marginLeft: 5
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
  return n._value/(SIZE/2);
}