import 'package:flutter/material.dart';

import '../../common/colors.dart';

class Dot extends StatelessWidget {

  final String dotTitle;
  final Color ?dotColor;
  final Color ?titleColor;


  const Dot({
    required this.dotTitle,
    this.dotColor,
    this.titleColor
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Container(
          height: 8,
          width: 8,
          margin: EdgeInsets.only(right: 9, top: 3),
          decoration: BoxDecoration(
              borderRadius: BorderRadius.all(Radius.circular(4)),
              color: dotColor ?? onlineColor),
        ),
        SelectableText(
          dotTitle,
          style: TextStyle(
            color: titleColor ?? regularFontColor,
            fontSize: 12,
            fontWeight: FontWeight.normal,
            fontFamily: 'Roboto',
            fontStyle: FontStyle.normal,
          )
        )
      ],
    );
  }
}
