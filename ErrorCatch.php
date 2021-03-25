<?php

/**
 *
 * 错误处理机制
 * @author wystanxu <wystanxu@tencent.com>
 * @since 1.0.0
 *
 */
class ErrorCatch
{
    /**
     * 注册异常处理
     * @access public
     * @return void
     */
    public static function register()
    {
        // 设定报错级别为全部
        error_reporting(E_ALL);
        // 设定错误日志记录文件
        ini_set('error_log', 'test_' . date("Ymd") . '.log');
        // set_error_handler — 设置用户自定义的错误处理函数
        set_error_handler([__CLASS__, 'appError']);
        // set_exception_handler — 设置用户自定义的异常处理函数
        set_exception_handler([__CLASS__, 'appException']);
        // register_shutdown_function — 注册一个会在php中止时执行的函数,脚本执行完成或者 exit() 后被调用
        register_shutdown_function([__CLASS__, 'appShutdown']);
    }

    /**
     * 错误处理
     * @access public
     * @param  integer $errno      错误编号
     * @param  integer $errstr     详细错误信息
     * @param  string  $errfile    出错的文件
     * @param  integer $errline    出错行号
     * @param  array   $errcontext 出错上下文
     * @return void
     * @throws ErrorException
     */
    public static function appError($errno, $errstr, $errfile = '', $errline = 0, $errcontext = [])
    {
        $msg  = date("Y-m-d H:i:s") . "\n\r";
        $msg .= '错误编号:' . $errno . "\n\r";
        $msg .= '错误信息:' . $errstr . "\n\r";
        $msg .= '文件:' . $errfile . "\n\r";
        $msg .= '在第:' . $errline . "行\n\r";
        echo $msg;
        // 记录错误日志到文本
        error_log($msg, 3, ini_get("error_log"));
        // 发送邮件通知
        self::errorReportDevelopment($errstr);
    }

    /**
     * 异常处理
     * @access public
     * @param   Exception $e 异常对象
     * @return void
     */
    public static function appException($exception)
    {
        $msg = "捕获异常: " . $exception->getMessage() . "\n\r";
        // 记录错误日志到文本
        error_log($msg, 3, ini_get("error_log"));

        echo $msg;
        // 发送邮件通知
        self::errorReportDevelopment($exception->getMessage());
        // 输出错误信息
        echo json_encode(array(
            'retcode' => "1",
            'data' => "",
            'msg' => "系统错误，请稍后重试或联系开发人员~"
        ));
        exit();
    }

    /**
     * 异常中止处理
     * @access public
     * @return void
     */
    public static function appShutdown()
    {
        // 只有错误导致的程序终止才会托管至错误处理函数
        if (!is_null($error = error_get_last()) && self::isFatal($error['type'])) {
            self::appException(new ErrorException(
                $error['type'],
                $error['message'],
                $error['file'],
                $error['line']
            ));
        }
    }

    /**
     * 确定错误类型是否致命
     * @access protected
     * @param  int $type 错误类型
     * @return bool
     */
    protected static function isFatal($type)
    {
        return in_array($type, [E_ERROR, E_CORE_ERROR, E_COMPILE_ERROR, E_PARSE]);
    }

    /**
     * 发送告警消息
     *
     * @param string $msg
     * @return void
     */
    protected static function errorReportDevelopment($msg)
    {
        // do something email ...
    }
}

// 注册自定义错误处理
ErrorCatch::register();
