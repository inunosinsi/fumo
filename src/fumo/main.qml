import QtQuick 2.4
import QtQuick.Controls 1.3
import QtQuick.Window 2.2
import QtQuick.Layouts 1.0
import QtQuick.Dialogs 1.2

ApplicationWindow {
    title: qsTr("Qt TextArea")
    width: 640
    height: 480
    visible: true
	
    FileDialog {
        id: fileDialog
        nameFilters: [ "Text files (*.txt)" ]
        onAccepted: sampleAction.insert(fileUrl)
    }
	
    //開くことに関するActionオブジェクト
    Action {
        id: openAction
        text: "&Open"
        shortcut: StandardKey.Open
        onTriggered: fileDialog.open()
//        tooltip: "Open an file"
    }
	
	Action {
		id: sampleAction
		function insert(source){
			textArea.text = "hoge"
		}
	}

    menuBar: MenuBar {
        Menu {
            title: qsTr("&File")
            MenuItem { action: openAction}
            MenuItem {
                text: qsTr("&Exit")
                onTriggered: Qt.quit();
            }
        }
        Menu {
            title: qsTr("&Help")
            MenuItem {
                text: qsTr("&Version")
            }
        }
    }

    TextArea{
        id: textArea
        anchors.bottom: parent.bottom
        y: partition
        width: parent.width
        height: parent.height
    }
}
