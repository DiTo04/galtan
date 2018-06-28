/* eslint-disable */
import React from 'react'
import {View, Animated, PanResponder, Image} from 'react-native'

const s_logo = require('./logos/s.png')

const styles = {
  container: {
    backgroundColor: '#FFFF0F',
    justifyContent: 'center',
    alignItems: 'center',
    flex: 1
  },
  square: {
    height: 64,
    width: 64
  }
}

class PanResponderExample extends React.Component {

  constructor(props) {
    super(props)

    this._dragAnimation = new Animated.ValueXY()

    this._onPanResponderGrant = this._onPanResponderGrant.bind(this)
    this._onPanResponderMove = this._onPanResponderMove.bind(this)
    this._onPanResponderRelease = this._onPanResponderRelease.bind(this)
    this._onPanResponderTerminate = this._onPanResponderTerminate.bind(this)
  }

  componentWillMount() {
    this._animatedValueX = 0
    this._animatedValueY = 0

    this._dragAnimation.x.addListener(({value}) => {
      this._animatedValueX = value
    })

    this._dragAnimation.y.addListener(({value}) => {
      this._animatedValueY = value
    })

    this._panResponder = PanResponder.create({
      // Ask to be the responder
      onStartShouldSetPanResponder: (evt, gestureState) => true,
      onStartShouldSetResponderCapture: (evt, gestureState) => true,
      onMoveShouldSetPanResponder: (evt, gestureState) => true,
      onMoveShouldSetResponderCapture: (evt, gestureState) => true,

      onPanResponderGrant: this._onPanResponderGrant,
      onPanResponderMove: this._onPanResponderMove,
      onPanResponderRelease: this._onPanResponderRelease,
      onPanResponderTerminate: this._onPanResponderTerminate
    })
  }

  _onPanResponderGrant(evt, gestureState) {
    this._dragAnimation.setOffset({
      x: this._animatedValueX, y: this._animatedValueY
    })

    this._dragAnimation.setValue({x: 0, y: 0})
  }

  _onPanResponderMove(evt, gestureState) {
    this._dragAnimation.setValue({
      x: gestureState.dx,
      y: gestureState.dy
    })
  }

  _onPanResponderRelease(evt, gestureState) {
    this._dragAnimation.flattenOffset()
  }

  _onPanResponderTerminate(evt, gestureState) {
    //
  }

  render() {
    const transform = this._dragAnimation.getTranslateTransform()

    return (
      <View style={styles.container}>
        <Animated.View
          {...this._panResponder.panHandlers}
          style={[styles.square, {transform}]}>
          <Image
          style={{width: 64, height: 64}}
          source={this.props.logo} />
        </Animated.View>
      </View>
    )
  }
}

export default PanResponderExample