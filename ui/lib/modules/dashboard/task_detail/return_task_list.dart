import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

import '../../../common/colors.dart';

class ReturnToList extends StatelessWidget {
  final String title;
  final Function onTap;

  const ReturnToList({
    required this.title,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(bottom: 10),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          MouseRegion(
            cursor: SystemMouseCursors.click,
            child: GestureDetector(
              child: Container(
                margin: EdgeInsets.only(top: 3, right: 8),
                child: Icon(IconData(0xe607, fontFamily: 'tpIcon'), size: 20,),
              ),
              onTap: () => onTap(),
            ),
          ),
          MouseRegion(
            cursor: SystemMouseCursors.click,
            child: GestureDetector(
              child: Text(
                title, 
                style: TextStyle(
                    fontSize: 14,
                    fontFamily: 'Roboto',
                    fontWeight: FontWeight.normal,
                    fontStyle: FontStyle.normal,
                    color: regularFontColor,
                ),
              ),
              onTap: () => onTap(),
            ),
          ),
        ],
      ),
    );
  }
}
