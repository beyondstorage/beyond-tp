import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';
import 'package:ui/common/svg_provider.dart';

class SvgDot extends StatelessWidget {
  final double ?size;
  final String src;
  final String dotTitle;
  final Color ?titleColor;

  const SvgDot({
    this.size = 16,
    required this.src,
    required this.dotTitle,
    this.titleColor
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Container(
          height: size,
          width: size,
          margin: EdgeInsets.only(top: 5, right: 9),
          decoration: BoxDecoration(
            image: DecorationImage(
              image: SvgProvider(
              src,
              size: Size(128, 128),
              color: rgba(255, 255, 255, 1),
            ),
            fit: BoxFit.fill,
          ),
          ),
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
