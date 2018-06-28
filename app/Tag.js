import React, { Component } from 'react';

import {
  StyleSheet,
  Text,
  View,
  Image, // we want to use an image
  PanResponder, // we want to bring in the PanResponder system
  Animated // we wil be using animated value
} from 'react-native';

export default class Tag extends React.Component {

constructor(props) {
  super(props);

  this.state = {
  pan: new Animated.ValueXY(),
  scale: new Animated.Value(1)
  };
}
_handleStartShouldSetPanResponder(e, gestureState) {
  return true;
}

_handleMoveShouldSetPanResponder(e, gestureState) {
 return true;
}

componentWillMount() {
  this._panResponder = PanResponder.create({
    onStartShouldSetPanResponder: 
  this._handleStartShouldSetPanResponder.bind(this),
    onMoveShouldSetPanResponder: 
  this._handleMoveShouldSetPanResponder.bind(this),

  onPanResponderGrant: (e, gestureState) => {
    // Set the initial value to the current state
    this.state.pan.setOffset({x: this.state.pan.x._value, y: this.state.pan.y._value});
    this.state.pan.setValue({x: 0, y: 0});
    Animated.spring(
      this.state.scale,
      { toValue: 1.1, friction: 1 }
    ).start();
  },

  // When we drag/pan the object, set the delate to the states pan position
  onPanResponderMove: Animated.event([
    null, {dx: this.state.pan.x, dy: this.state.pan.y},
  ]),

  onPanResponderRelease: (e, {vx, vy}) => {
    let obj = {};
    obj[this.props.party] = this.state.pan;
    this.props.updateParty(obj, this.props.party);
    // Flatten the offset to avoid erratic behavior
    this.state.pan.flattenOffset();
    Animated.spring(
      this.state.scale,
      { toValue: 1, friction: 1 }
    ).start();
    }
   });
  }

  render() {
// Destructure the value of pan from the state
let { pan, scale } = this.state;

// Calculate the x and y transform from the pan value
let [translateX, translateY] = [pan.x, pan.y];

let rotate = '0deg';
// Calculate the transform property and set it as a value for our style which we add below to the Animated.View component
let imageStyle = {transform: [{translateX}, {translateY}, {rotate}, {scale}]};


return (
    <Animated.View style={[imageStyle, styles.container]} {...this._panResponder.panHandlers} >
      <View style={styles.rect}>
        <Image
          style={{width: 96, height: 96}}
          source={this.props.logo} />
      </View>
    </Animated.View>
);
}

}

const styles = StyleSheet.create({
  container: {
  width:96,
  height:96,
  position: 'absolute'
},
rect: {
  borderRadius:4,
  borderWidth: 1,
  borderColor: '#fff',
  width:96,
  height:96

  },
  txt: {
    color:'#fff',
    textAlign:'center'
  }

 });