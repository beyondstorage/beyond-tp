import 'package:flutter/material.dart';

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
      child: GestureDetector(
        child: Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Container(
              margin: EdgeInsets.only(top: 3, right: 8),
              child: Icon(IconData(0xe607, fontFamily: 'tpIcon'), size: 20,),
            ),
            SelectableText(
              title, 
              style: TextStyle(
                  fontSize: 14,
                  fontFamily: 'Roboto',
                  fontWeight: FontWeight.normal,
                  fontStyle: FontStyle.normal,
                  color: regularFontColor,
              ),
            ),
          ],
        ),
        onTap: () => onTap(),
      ),
    );
  }
}
