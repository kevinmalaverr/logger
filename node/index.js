const fs = require('fs');

const colors = {
  success: '\x1b[32m',
  info: '\x1b[34m',
  warn: '\x1b[33m',
  error: '\x1b[31m',
  default: '\x1b[0m',
};

const getFormatDate = () => {
  return new Date().toISOString().replace(/T/, ' ').replace(/\..+/, '');
};

function tabAlign(str, length = 10) {
  return str + ' '.repeat(length - str.length);
}

class Logger {
  static logPath = '.log/';
  static logNames = {
    success: 'SUCCESS',
    info: 'INFO',
    warn: 'WARN',
    error: 'ERROR',
  };

  static measure(message, fun) {
    console.time(message);
    fun();
    console.timeEnd(message);
  }

  static writer;

  static initWriteLog() {
    if (!fs.existsSync(this.logPath)) {
      fs.mkdirSync(this.logPath);
    }

    const time = getFormatDate();
    const fileName = this.logPath + time + '.log';

    fs.writeFileSync(fileName, '');

    this.writer = fs.createWriteStream(fileName, { flags: 'a' });
  }

  static writeLog(message) {
    if (this.writer) this.writer.write(`${getFormatDate()} ${message}\n`);
  }

  static log(type, message) {
    console.log(
      `${colors[type]}${this.logNames[type]}${colors.default}`,
      ...message
    );
    this.writeLog(`${tabAlign(this.logNames[type])} ${message.join(' ')}`);
  }

  static success(...message) {
    this.log('success', message);
  }

  static info(...message) {
    this.log('info', message);
  }

  static warn(...message) {
    this.log('warn', message);
  }

  static error(...message) {
    this.log('error', message);
  }
}

module.exports = Logger;
