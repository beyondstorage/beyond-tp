import 'package:flutter/material.dart';
import 'package:ui/widgets/message/animation.dart';

import '../../common/colors.dart';

class Message {
  static void success({required BuildContext context, required String message}) {
    late OverlayEntry overlayEntry;
    //cretae an OverlayEntry object
    overlayEntry = new OverlayEntry(builder: (context) {
      return MessageAnimation(
        message: message,
        fontColor: successFontColor,
        icon: Icon(
          Icons.check_circle,
          size: 22,
          color: successFontColor,
        ),
        backgroundColor: successMessageColor,
        callBack: () {
          if (overlayEntry != null) {
            overlayEntry.remove();
          }
        }
      );
    });
    Overlay.of(context, rootOverlay: true)?.insert(overlayEntry);
  }
}
