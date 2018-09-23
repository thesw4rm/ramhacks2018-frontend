package cs.vcu.edu.cmsc355.ramdev_2018

import android.annotation.SuppressLint
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.view.Window
import android.view.WindowManager
import android.webkit.WebView

class MainActivity : AppCompatActivity() {

    @SuppressLint("SetJavaScriptEnabled")
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        requestWindowFeature(Window.FEATURE_NO_TITLE)
        window.setFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN,
                WindowManager.LayoutParams.FLAG_FULLSCREEN)
        val myWebView = WebView(applicationContext)
        myWebView.settings.javaScriptEnabled = true
        myWebView.clearCache(true)
        setContentView(myWebView)
        myWebView.loadUrl("https://helphunter.tk/")
    }
}
