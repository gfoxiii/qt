package webengine

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_NETWORK_LIB -DQT_WEBCHANNEL_LIB -DQT_WIDGETS_LIB -DQT_WEBENGINECORE_LIB -DQT_WEBENGINEWIDGETS_LIB -DQT_WEBENGINE_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc_64/include -I/usr/local/Qt5.6.0/5.6/gcc_64/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtCore -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtGui -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtNetwork -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtWebChannel -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtWidgets -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtWebEngineCore -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtWebEngineWidgets -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtWebEngine

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc_64 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc_64/lib
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc_64/lib -L/usr/lib64 -lQt5Core -lQt5Gui -lQt5Network -lQt5WebChannel -lQt5Widgets -lQt5WebEngineCore -lQt5WebEngineWidgets -lQt5WebEngine -lpthread
*/
import "C"