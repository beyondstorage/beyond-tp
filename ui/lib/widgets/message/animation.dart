import 'package:flutter/material.dart';

import '../../common/colors.dart';

class MessageAnimation extends StatefulWidget {
  final String message;
  final Color backgroundColor;
  final Icon icon;
  final Color fontColor;
  final Function callBack;
  final int duration;
  const MessageAnimation({
    required this.message,
    required this.backgroundColor,
    required this.icon,
    required this.fontColor,
    required this.callBack,
    this.duration = 1500,
  });
  @override
  _MessageAnimationState createState() => _MessageAnimationState();
}

class _MessageAnimationState extends State<MessageAnimation> with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  late Animation<double> animation;
  double fadeValue = 0;
  double marginValue = 10;
  @override
  void initState() {
    super.initState();
    _controller = AnimationController(vsync: this, duration: Duration(milliseconds: 500));
    animation = Tween<double>(begin: 0, end: 1).animate(_controller)
      ..addListener(() {
        setState(() {
          fadeValue = animation.value;
          marginValue = 10 + animation.value * 20;
        });
      })
      ..addStatusListener((status) {
        if (status == AnimationStatus.completed) {
          new Future.delayed(Duration(milliseconds: widget.duration)).then((value) {
            _controller.reverse();
          });
        } else if (status == AnimationStatus.dismissed) {
          widget.callBack();
        }
      });
    _controller.reset();
    _controller.forward();
  }

  @override
  void dispose() {
    super.dispose();
    _controller.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // todo The first time it was used, it got stuck
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        Opacity(
          opacity: fadeValue,
          child: Container(
            margin: EdgeInsets.only(top: marginValue),
            constraints:
                BoxConstraints(minWidth: 120, maxWidth: 240, minHeight: 40),
            alignment: Alignment.center,
            decoration: BoxDecoration(
                color: widget.backgroundColor,
                border: Border.all(width: 1, color: widget.fontColor),
                borderRadius: BorderRadius.all(Radius.circular(4))),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Container(
                  margin: EdgeInsets.symmetric(horizontal: 10),
                  child: widget.icon,
                ),
                Text(
                  widget.message,
                  style: TextStyle(
                    fontSize: 16,
                    fontWeight: FontWeight.w500,
                    color: widget.fontColor,
                    decoration: TextDecoration.none,
                  ),
                )
              ],
            ),
          ),
        )
      ],
    );
  }
}
