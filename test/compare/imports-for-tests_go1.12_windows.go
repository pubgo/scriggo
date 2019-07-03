// Code generated by scriggo-generate, based on file "imports-for-tests.go". DO NOT EDIT.
//+build windows,go1.12,!go1.13

package main

import (
	bufio "bufio"
	bytes "bytes"
	sha1 "crypto/sha1"
	base64 "encoding/base64"
	fmt "fmt"
	ioutil "io/ioutil"
	math "math"
	rand "math/rand"
	net "net"
	url "net/url"
	os "os"
	regexp "regexp"
	sort "sort"
	strconv "strconv"
	strings "strings"
	time "time"
)

import . "scriggo"
import "reflect"

func init() {
	packages = Packages{

		"io/ioutil": {
			Name: "ioutil",
			Declarations: map[string]interface{}{
				"Discard":   &ioutil.Discard,
				"NopCloser": ioutil.NopCloser,
				"ReadAll":   ioutil.ReadAll,
				"ReadDir":   ioutil.ReadDir,
				"ReadFile":  ioutil.ReadFile,
				"TempDir":   ioutil.TempDir,
				"TempFile":  ioutil.TempFile,
				"WriteFile": ioutil.WriteFile,
			},
		},
		"net": {
			Name: "net",
			Declarations: map[string]interface{}{
				"Addr":                       reflect.TypeOf(new(net.Addr)).Elem(),
				"AddrError":                  reflect.TypeOf(net.AddrError{}),
				"Buffers":                    reflect.TypeOf(new(net.Buffers)).Elem(),
				"CIDRMask":                   net.CIDRMask,
				"Conn":                       reflect.TypeOf(new(net.Conn)).Elem(),
				"DNSConfigError":             reflect.TypeOf(net.DNSConfigError{}),
				"DNSError":                   reflect.TypeOf(net.DNSError{}),
				"DefaultResolver":            &net.DefaultResolver,
				"Dial":                       net.Dial,
				"DialIP":                     net.DialIP,
				"DialTCP":                    net.DialTCP,
				"DialTimeout":                net.DialTimeout,
				"DialUDP":                    net.DialUDP,
				"DialUnix":                   net.DialUnix,
				"Dialer":                     reflect.TypeOf(net.Dialer{}),
				"ErrWriteToConnected":        &net.ErrWriteToConnected,
				"Error":                      reflect.TypeOf(new(net.Error)).Elem(),
				"FileConn":                   net.FileConn,
				"FileListener":               net.FileListener,
				"FilePacketConn":             net.FilePacketConn,
				"FlagBroadcast":              ConstValue(net.FlagBroadcast),
				"FlagLoopback":               ConstValue(net.FlagLoopback),
				"FlagMulticast":              ConstValue(net.FlagMulticast),
				"FlagPointToPoint":           ConstValue(net.FlagPointToPoint),
				"FlagUp":                     ConstValue(net.FlagUp),
				"Flags":                      reflect.TypeOf(new(net.Flags)).Elem(),
				"HardwareAddr":               reflect.TypeOf(new(net.HardwareAddr)).Elem(),
				"IP":                         reflect.TypeOf(new(net.IP)).Elem(),
				"IPAddr":                     reflect.TypeOf(net.IPAddr{}),
				"IPConn":                     reflect.TypeOf(net.IPConn{}),
				"IPMask":                     reflect.TypeOf(new(net.IPMask)).Elem(),
				"IPNet":                      reflect.TypeOf(net.IPNet{}),
				"IPv4":                       net.IPv4,
				"IPv4Mask":                   net.IPv4Mask,
				"IPv4allrouter":              &net.IPv4allrouter,
				"IPv4allsys":                 &net.IPv4allsys,
				"IPv4bcast":                  &net.IPv4bcast,
				"IPv4len":                    ConstValue(net.IPv4len),
				"IPv4zero":                   &net.IPv4zero,
				"IPv6interfacelocalallnodes": &net.IPv6interfacelocalallnodes,
				"IPv6len":                    ConstValue(net.IPv6len),
				"IPv6linklocalallnodes":      &net.IPv6linklocalallnodes,
				"IPv6linklocalallrouters":    &net.IPv6linklocalallrouters,
				"IPv6loopback":               &net.IPv6loopback,
				"IPv6unspecified":            &net.IPv6unspecified,
				"IPv6zero":                   &net.IPv6zero,
				"Interface":                  reflect.TypeOf(net.Interface{}),
				"InterfaceAddrs":             net.InterfaceAddrs,
				"InterfaceByIndex":           net.InterfaceByIndex,
				"InterfaceByName":            net.InterfaceByName,
				"Interfaces":                 net.Interfaces,
				"InvalidAddrError":           reflect.TypeOf(new(net.InvalidAddrError)).Elem(),
				"JoinHostPort":               net.JoinHostPort,
				"Listen":                     net.Listen,
				"ListenConfig":               reflect.TypeOf(net.ListenConfig{}),
				"ListenIP":                   net.ListenIP,
				"ListenMulticastUDP":         net.ListenMulticastUDP,
				"ListenPacket":               net.ListenPacket,
				"ListenTCP":                  net.ListenTCP,
				"ListenUDP":                  net.ListenUDP,
				"ListenUnix":                 net.ListenUnix,
				"ListenUnixgram":             net.ListenUnixgram,
				"Listener":                   reflect.TypeOf(new(net.Listener)).Elem(),
				"LookupAddr":                 net.LookupAddr,
				"LookupCNAME":                net.LookupCNAME,
				"LookupHost":                 net.LookupHost,
				"LookupIP":                   net.LookupIP,
				"LookupMX":                   net.LookupMX,
				"LookupNS":                   net.LookupNS,
				"LookupPort":                 net.LookupPort,
				"LookupSRV":                  net.LookupSRV,
				"LookupTXT":                  net.LookupTXT,
				"MX":                         reflect.TypeOf(net.MX{}),
				"NS":                         reflect.TypeOf(net.NS{}),
				"OpError":                    reflect.TypeOf(net.OpError{}),
				"PacketConn":                 reflect.TypeOf(new(net.PacketConn)).Elem(),
				"ParseCIDR":                  net.ParseCIDR,
				"ParseError":                 reflect.TypeOf(net.ParseError{}),
				"ParseIP":                    net.ParseIP,
				"ParseMAC":                   net.ParseMAC,
				"Pipe":                       net.Pipe,
				"ResolveIPAddr":              net.ResolveIPAddr,
				"ResolveTCPAddr":             net.ResolveTCPAddr,
				"ResolveUDPAddr":             net.ResolveUDPAddr,
				"ResolveUnixAddr":            net.ResolveUnixAddr,
				"Resolver":                   reflect.TypeOf(net.Resolver{}),
				"SRV":                        reflect.TypeOf(net.SRV{}),
				"SplitHostPort":              net.SplitHostPort,
				"TCPAddr":                    reflect.TypeOf(net.TCPAddr{}),
				"TCPConn":                    reflect.TypeOf(net.TCPConn{}),
				"TCPListener":                reflect.TypeOf(net.TCPListener{}),
				"UDPAddr":                    reflect.TypeOf(net.UDPAddr{}),
				"UDPConn":                    reflect.TypeOf(net.UDPConn{}),
				"UnixAddr":                   reflect.TypeOf(net.UnixAddr{}),
				"UnixConn":                   reflect.TypeOf(net.UnixConn{}),
				"UnixListener":               reflect.TypeOf(net.UnixListener{}),
				"UnknownNetworkError":        reflect.TypeOf(new(net.UnknownNetworkError)).Elem(),
			},
		},
		"os": {
			Name: "os",
			Declarations: map[string]interface{}{
				"Args":              &os.Args,
				"Chdir":             os.Chdir,
				"Chmod":             os.Chmod,
				"Chown":             os.Chown,
				"Chtimes":           os.Chtimes,
				"Clearenv":          os.Clearenv,
				"Create":            os.Create,
				"DevNull":           ConstValue(os.DevNull),
				"Environ":           os.Environ,
				"ErrClosed":         &os.ErrClosed,
				"ErrExist":          &os.ErrExist,
				"ErrInvalid":        &os.ErrInvalid,
				"ErrNoDeadline":     &os.ErrNoDeadline,
				"ErrNotExist":       &os.ErrNotExist,
				"ErrPermission":     &os.ErrPermission,
				"Executable":        os.Executable,
				"Exit":              os.Exit,
				"Expand":            os.Expand,
				"ExpandEnv":         os.ExpandEnv,
				"File":              reflect.TypeOf(os.File{}),
				"FileInfo":          reflect.TypeOf(new(os.FileInfo)).Elem(),
				"FileMode":          reflect.TypeOf(new(os.FileMode)).Elem(),
				"FindProcess":       os.FindProcess,
				"Getegid":           os.Getegid,
				"Getenv":            os.Getenv,
				"Geteuid":           os.Geteuid,
				"Getgid":            os.Getgid,
				"Getgroups":         os.Getgroups,
				"Getpagesize":       os.Getpagesize,
				"Getpid":            os.Getpid,
				"Getppid":           os.Getppid,
				"Getuid":            os.Getuid,
				"Getwd":             os.Getwd,
				"Hostname":          os.Hostname,
				"Interrupt":         &os.Interrupt,
				"IsExist":           os.IsExist,
				"IsNotExist":        os.IsNotExist,
				"IsPathSeparator":   os.IsPathSeparator,
				"IsPermission":      os.IsPermission,
				"IsTimeout":         os.IsTimeout,
				"Kill":              &os.Kill,
				"Lchown":            os.Lchown,
				"Link":              os.Link,
				"LinkError":         reflect.TypeOf(os.LinkError{}),
				"LookupEnv":         os.LookupEnv,
				"Lstat":             os.Lstat,
				"Mkdir":             os.Mkdir,
				"MkdirAll":          os.MkdirAll,
				"ModeAppend":        ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "1073741824"),
				"ModeCharDevice":    ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "2097152"),
				"ModeDevice":        ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "67108864"),
				"ModeDir":           ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "2147483648"),
				"ModeExclusive":     ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "536870912"),
				"ModeIrregular":     ConstValue(os.ModeIrregular),
				"ModeNamedPipe":     ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "33554432"),
				"ModePerm":          ConstValue(os.ModePerm),
				"ModeSetgid":        ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "4194304"),
				"ModeSetuid":        ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "8388608"),
				"ModeSocket":        ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "16777216"),
				"ModeSticky":        ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "1048576"),
				"ModeSymlink":       ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "134217728"),
				"ModeTemporary":     ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "268435456"),
				"ModeType":          ConstLiteral(reflect.TypeOf(new(os.FileMode)).Elem(), "2401763328"),
				"NewFile":           os.NewFile,
				"NewSyscallError":   os.NewSyscallError,
				"O_APPEND":          ConstValue(os.O_APPEND),
				"O_CREATE":          ConstValue(os.O_CREATE),
				"O_EXCL":            ConstValue(os.O_EXCL),
				"O_RDONLY":          ConstValue(os.O_RDONLY),
				"O_RDWR":            ConstValue(os.O_RDWR),
				"O_SYNC":            ConstValue(os.O_SYNC),
				"O_TRUNC":           ConstValue(os.O_TRUNC),
				"O_WRONLY":          ConstValue(os.O_WRONLY),
				"Open":              os.Open,
				"OpenFile":          os.OpenFile,
				"PathError":         reflect.TypeOf(os.PathError{}),
				"PathListSeparator": ConstValue(os.PathListSeparator),
				"PathSeparator":     ConstValue(os.PathSeparator),
				"Pipe":              os.Pipe,
				"ProcAttr":          reflect.TypeOf(os.ProcAttr{}),
				"Process":           reflect.TypeOf(os.Process{}),
				"ProcessState":      reflect.TypeOf(os.ProcessState{}),
				"Readlink":          os.Readlink,
				"Remove":            os.Remove,
				"RemoveAll":         os.RemoveAll,
				"Rename":            os.Rename,
				"SEEK_CUR":          ConstValue(os.SEEK_CUR),
				"SEEK_END":          ConstValue(os.SEEK_END),
				"SEEK_SET":          ConstValue(os.SEEK_SET),
				"SameFile":          os.SameFile,
				"Setenv":            os.Setenv,
				"Signal":            reflect.TypeOf(new(os.Signal)).Elem(),
				"StartProcess":      os.StartProcess,
				"Stat":              os.Stat,
				"Stderr":            &os.Stderr,
				"Stdin":             &os.Stdin,
				"Stdout":            &os.Stdout,
				"Symlink":           os.Symlink,
				"SyscallError":      reflect.TypeOf(os.SyscallError{}),
				"TempDir":           os.TempDir,
				"Truncate":          os.Truncate,
				"Unsetenv":          os.Unsetenv,
				"UserCacheDir":      os.UserCacheDir,
				"UserHomeDir":       os.UserHomeDir,
			},
		},
		"encoding/base64": {
			Name: "base64",
			Declarations: map[string]interface{}{
				"CorruptInputError": reflect.TypeOf(new(base64.CorruptInputError)).Elem(),
				"Encoding":          reflect.TypeOf(base64.Encoding{}),
				"NewDecoder":        base64.NewDecoder,
				"NewEncoder":        base64.NewEncoder,
				"NewEncoding":       base64.NewEncoding,
				"NoPadding":         ConstValue(base64.NoPadding),
				"RawStdEncoding":    &base64.RawStdEncoding,
				"RawURLEncoding":    &base64.RawURLEncoding,
				"StdEncoding":       &base64.StdEncoding,
				"StdPadding":        ConstValue(base64.StdPadding),
				"URLEncoding":       &base64.URLEncoding,
			},
		},
		"fmt": {
			Name: "fmt",
			Declarations: map[string]interface{}{
				"Errorf":     fmt.Errorf,
				"Formatter":  reflect.TypeOf(new(fmt.Formatter)).Elem(),
				"Fprint":     fmt.Fprint,
				"Fprintf":    fmt.Fprintf,
				"Fprintln":   fmt.Fprintln,
				"Fscan":      fmt.Fscan,
				"Fscanf":     fmt.Fscanf,
				"Fscanln":    fmt.Fscanln,
				"GoStringer": reflect.TypeOf(new(fmt.GoStringer)).Elem(),
				"Print":      fmt.Print,
				"Printf":     fmt.Printf,
				"Println":    fmt.Println,
				"Scan":       fmt.Scan,
				"ScanState":  reflect.TypeOf(new(fmt.ScanState)).Elem(),
				"Scanf":      fmt.Scanf,
				"Scanln":     fmt.Scanln,
				"Scanner":    reflect.TypeOf(new(fmt.Scanner)).Elem(),
				"Sprint":     fmt.Sprint,
				"Sprintf":    fmt.Sprintf,
				"Sprintln":   fmt.Sprintln,
				"Sscan":      fmt.Sscan,
				"Sscanf":     fmt.Sscanf,
				"Sscanln":    fmt.Sscanln,
				"State":      reflect.TypeOf(new(fmt.State)).Elem(),
				"Stringer":   reflect.TypeOf(new(fmt.Stringer)).Elem(),
			},
		},
		"net/url": {
			Name: "url",
			Declarations: map[string]interface{}{
				"Error":            reflect.TypeOf(url.Error{}),
				"EscapeError":      reflect.TypeOf(new(url.EscapeError)).Elem(),
				"InvalidHostError": reflect.TypeOf(new(url.InvalidHostError)).Elem(),
				"Parse":            url.Parse,
				"ParseQuery":       url.ParseQuery,
				"ParseRequestURI":  url.ParseRequestURI,
				"PathEscape":       url.PathEscape,
				"PathUnescape":     url.PathUnescape,
				"QueryEscape":      url.QueryEscape,
				"QueryUnescape":    url.QueryUnescape,
				"URL":              reflect.TypeOf(url.URL{}),
				"User":             url.User,
				"UserPassword":     url.UserPassword,
				"Userinfo":         reflect.TypeOf(url.Userinfo{}),
				"Values":           reflect.TypeOf(new(url.Values)).Elem(),
			},
		},
		"regexp": {
			Name: "regexp",
			Declarations: map[string]interface{}{
				"Compile":          regexp.Compile,
				"CompilePOSIX":     regexp.CompilePOSIX,
				"Match":            regexp.Match,
				"MatchReader":      regexp.MatchReader,
				"MatchString":      regexp.MatchString,
				"MustCompile":      regexp.MustCompile,
				"MustCompilePOSIX": regexp.MustCompilePOSIX,
				"QuoteMeta":        regexp.QuoteMeta,
				"Regexp":           reflect.TypeOf(regexp.Regexp{}),
			},
		},
		"bufio": {
			Name: "bufio",
			Declarations: map[string]interface{}{
				"ErrAdvanceTooFar":     &bufio.ErrAdvanceTooFar,
				"ErrBufferFull":        &bufio.ErrBufferFull,
				"ErrFinalToken":        &bufio.ErrFinalToken,
				"ErrInvalidUnreadByte": &bufio.ErrInvalidUnreadByte,
				"ErrInvalidUnreadRune": &bufio.ErrInvalidUnreadRune,
				"ErrNegativeAdvance":   &bufio.ErrNegativeAdvance,
				"ErrNegativeCount":     &bufio.ErrNegativeCount,
				"ErrTooLong":           &bufio.ErrTooLong,
				"MaxScanTokenSize":     ConstValue(bufio.MaxScanTokenSize),
				"NewReadWriter":        bufio.NewReadWriter,
				"NewReader":            bufio.NewReader,
				"NewReaderSize":        bufio.NewReaderSize,
				"NewScanner":           bufio.NewScanner,
				"NewWriter":            bufio.NewWriter,
				"NewWriterSize":        bufio.NewWriterSize,
				"ReadWriter":           reflect.TypeOf(bufio.ReadWriter{}),
				"Reader":               reflect.TypeOf(bufio.Reader{}),
				"ScanBytes":            bufio.ScanBytes,
				"ScanLines":            bufio.ScanLines,
				"ScanRunes":            bufio.ScanRunes,
				"ScanWords":            bufio.ScanWords,
				"Scanner":              reflect.TypeOf(bufio.Scanner{}),
				"SplitFunc":            reflect.TypeOf(new(bufio.SplitFunc)).Elem(),
				"Writer":               reflect.TypeOf(bufio.Writer{}),
			},
		},
		"math": {
			Name: "math",
			Declarations: map[string]interface{}{
				"Abs":                    math.Abs,
				"Acos":                   math.Acos,
				"Acosh":                  math.Acosh,
				"Asin":                   math.Asin,
				"Asinh":                  math.Asinh,
				"Atan":                   math.Atan,
				"Atan2":                  math.Atan2,
				"Atanh":                  math.Atanh,
				"Cbrt":                   math.Cbrt,
				"Ceil":                   math.Ceil,
				"Copysign":               math.Copysign,
				"Cos":                    math.Cos,
				"Cosh":                   math.Cosh,
				"Dim":                    math.Dim,
				"E":                      ConstLiteral(nil, "271828182845904523536028747135266249775724709369995957496696763/100000000000000000000000000000000000000000000000000000000000000"),
				"Erf":                    math.Erf,
				"Erfc":                   math.Erfc,
				"Erfcinv":                math.Erfcinv,
				"Erfinv":                 math.Erfinv,
				"Exp":                    math.Exp,
				"Exp2":                   math.Exp2,
				"Expm1":                  math.Expm1,
				"Float32bits":            math.Float32bits,
				"Float32frombits":        math.Float32frombits,
				"Float64bits":            math.Float64bits,
				"Float64frombits":        math.Float64frombits,
				"Floor":                  math.Floor,
				"Frexp":                  math.Frexp,
				"Gamma":                  math.Gamma,
				"Hypot":                  math.Hypot,
				"Ilogb":                  math.Ilogb,
				"Inf":                    math.Inf,
				"IsInf":                  math.IsInf,
				"IsNaN":                  math.IsNaN,
				"J0":                     math.J0,
				"J1":                     math.J1,
				"Jn":                     math.Jn,
				"Ldexp":                  math.Ldexp,
				"Lgamma":                 math.Lgamma,
				"Ln10":                   ConstLiteral(nil, "23025850929940456840179914546843642076011014886287729760333279/10000000000000000000000000000000000000000000000000000000000000"),
				"Ln2":                    ConstLiteral(nil, "693147180559945309417232121458176568075500134360255254120680009/1000000000000000000000000000000000000000000000000000000000000000"),
				"Log":                    math.Log,
				"Log10":                  math.Log10,
				"Log10E":                 ConstLiteral(nil, "10000000000000000000000000000000000000000000000000000000000000/23025850929940456840179914546843642076011014886287729760333279"),
				"Log1p":                  math.Log1p,
				"Log2":                   math.Log2,
				"Log2E":                  ConstLiteral(nil, "1000000000000000000000000000000000000000000000000000000000000000/693147180559945309417232121458176568075500134360255254120680009"),
				"Logb":                   math.Logb,
				"Max":                    math.Max,
				"MaxFloat32":             ConstLiteral(nil, "340282346638528859811704183484516925440"),
				"MaxFloat64":             ConstLiteral(nil, "179769313486231570814527423731704356798100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				"MaxInt16":               ConstValue(math.MaxInt16),
				"MaxInt32":               ConstLiteral(nil, "2147483647"),
				"MaxInt64":               ConstLiteral(nil, "9223372036854775807"),
				"MaxInt8":                ConstValue(math.MaxInt8),
				"MaxUint16":              ConstValue(math.MaxUint16),
				"MaxUint32":              ConstLiteral(nil, "4294967295"),
				"MaxUint64":              ConstLiteral(nil, "18446744073709551615"),
				"MaxUint8":               ConstValue(math.MaxUint8),
				"Min":                    math.Min,
				"MinInt16":               ConstValue(math.MinInt16),
				"MinInt32":               ConstLiteral(nil, "-2147483648"),
				"MinInt64":               ConstLiteral(nil, "-9223372036854775808"),
				"MinInt8":                ConstValue(math.MinInt8),
				"Mod":                    math.Mod,
				"Modf":                   math.Modf,
				"NaN":                    math.NaN,
				"Nextafter":              math.Nextafter,
				"Nextafter32":            math.Nextafter32,
				"Phi":                    ConstLiteral(nil, "80901699437494742410229341718281905886015458990288143106772431/50000000000000000000000000000000000000000000000000000000000000"),
				"Pi":                     ConstLiteral(nil, "314159265358979323846264338327950288419716939937510582097494459/100000000000000000000000000000000000000000000000000000000000000"),
				"Pow":                    math.Pow,
				"Pow10":                  math.Pow10,
				"Remainder":              math.Remainder,
				"Round":                  math.Round,
				"RoundToEven":            math.RoundToEven,
				"Signbit":                math.Signbit,
				"Sin":                    math.Sin,
				"Sincos":                 math.Sincos,
				"Sinh":                   math.Sinh,
				"SmallestNonzeroFloat32": ConstLiteral(nil, "17516230804060213386546619791123951641/12500000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				"SmallestNonzeroFloat64": ConstLiteral(nil, "4940656458412465441765687928682213723651/1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
				"Sqrt":                   math.Sqrt,
				"Sqrt2":                  ConstLiteral(nil, "70710678118654752440084436210484903928483593768847403658833987/50000000000000000000000000000000000000000000000000000000000000"),
				"SqrtE":                  ConstLiteral(nil, "164872127070012814684865078781416357165377610071014801157507931/100000000000000000000000000000000000000000000000000000000000000"),
				"SqrtPhi":                ConstLiteral(nil, "63600982475703448212621123086874574585780402092004812430832019/50000000000000000000000000000000000000000000000000000000000000"),
				"SqrtPi":                 ConstLiteral(nil, "177245385090551602729816748334114518279754945612238712821380779/100000000000000000000000000000000000000000000000000000000000000"),
				"Tan":                    math.Tan,
				"Tanh":                   math.Tanh,
				"Trunc":                  math.Trunc,
				"Y0":                     math.Y0,
				"Y1":                     math.Y1,
				"Yn":                     math.Yn,
			},
		},
		"math/rand": {
			Name: "rand",
			Declarations: map[string]interface{}{
				"ExpFloat64":  rand.ExpFloat64,
				"Float32":     rand.Float32,
				"Float64":     rand.Float64,
				"Int":         rand.Int,
				"Int31":       rand.Int31,
				"Int31n":      rand.Int31n,
				"Int63":       rand.Int63,
				"Int63n":      rand.Int63n,
				"Intn":        rand.Intn,
				"New":         rand.New,
				"NewSource":   rand.NewSource,
				"NewZipf":     rand.NewZipf,
				"NormFloat64": rand.NormFloat64,
				"Perm":        rand.Perm,
				"Rand":        reflect.TypeOf(rand.Rand{}),
				"Read":        rand.Read,
				"Seed":        rand.Seed,
				"Shuffle":     rand.Shuffle,
				"Source":      reflect.TypeOf(new(rand.Source)).Elem(),
				"Source64":    reflect.TypeOf(new(rand.Source64)).Elem(),
				"Uint32":      rand.Uint32,
				"Uint64":      rand.Uint64,
				"Zipf":        reflect.TypeOf(rand.Zipf{}),
			},
		},
		"sort": {
			Name: "sort",
			Declarations: map[string]interface{}{
				"Float64Slice":      reflect.TypeOf(new(sort.Float64Slice)).Elem(),
				"Float64s":          sort.Float64s,
				"Float64sAreSorted": sort.Float64sAreSorted,
				"IntSlice":          reflect.TypeOf(new(sort.IntSlice)).Elem(),
				"Interface":         reflect.TypeOf(new(sort.Interface)).Elem(),
				"Ints":              sort.Ints,
				"IntsAreSorted":     sort.IntsAreSorted,
				"IsSorted":          sort.IsSorted,
				"Reverse":           sort.Reverse,
				"Search":            sort.Search,
				"SearchFloat64s":    sort.SearchFloat64s,
				"SearchInts":        sort.SearchInts,
				"SearchStrings":     sort.SearchStrings,
				"Slice":             sort.Slice,
				"SliceIsSorted":     sort.SliceIsSorted,
				"SliceStable":       sort.SliceStable,
				"Sort":              sort.Sort,
				"Stable":            sort.Stable,
				"StringSlice":       reflect.TypeOf(new(sort.StringSlice)).Elem(),
				"Strings":           sort.Strings,
				"StringsAreSorted":  sort.StringsAreSorted,
			},
		},
		"strconv": {
			Name: "strconv",
			Declarations: map[string]interface{}{
				"AppendBool":               strconv.AppendBool,
				"AppendFloat":              strconv.AppendFloat,
				"AppendInt":                strconv.AppendInt,
				"AppendQuote":              strconv.AppendQuote,
				"AppendQuoteRune":          strconv.AppendQuoteRune,
				"AppendQuoteRuneToASCII":   strconv.AppendQuoteRuneToASCII,
				"AppendQuoteRuneToGraphic": strconv.AppendQuoteRuneToGraphic,
				"AppendQuoteToASCII":       strconv.AppendQuoteToASCII,
				"AppendQuoteToGraphic":     strconv.AppendQuoteToGraphic,
				"AppendUint":               strconv.AppendUint,
				"Atoi":                     strconv.Atoi,
				"CanBackquote":             strconv.CanBackquote,
				"ErrRange":                 &strconv.ErrRange,
				"ErrSyntax":                &strconv.ErrSyntax,
				"FormatBool":               strconv.FormatBool,
				"FormatFloat":              strconv.FormatFloat,
				"FormatInt":                strconv.FormatInt,
				"FormatUint":               strconv.FormatUint,
				"IntSize":                  ConstValue(strconv.IntSize),
				"IsGraphic":                strconv.IsGraphic,
				"IsPrint":                  strconv.IsPrint,
				"Itoa":                     strconv.Itoa,
				"NumError":                 reflect.TypeOf(strconv.NumError{}),
				"ParseBool":                strconv.ParseBool,
				"ParseFloat":               strconv.ParseFloat,
				"ParseInt":                 strconv.ParseInt,
				"ParseUint":                strconv.ParseUint,
				"Quote":                    strconv.Quote,
				"QuoteRune":                strconv.QuoteRune,
				"QuoteRuneToASCII":         strconv.QuoteRuneToASCII,
				"QuoteRuneToGraphic":       strconv.QuoteRuneToGraphic,
				"QuoteToASCII":             strconv.QuoteToASCII,
				"QuoteToGraphic":           strconv.QuoteToGraphic,
				"Unquote":                  strconv.Unquote,
				"UnquoteChar":              strconv.UnquoteChar,
			},
		},
		"strings": {
			Name: "strings",
			Declarations: map[string]interface{}{
				"Builder":        reflect.TypeOf(strings.Builder{}),
				"Compare":        strings.Compare,
				"Contains":       strings.Contains,
				"ContainsAny":    strings.ContainsAny,
				"ContainsRune":   strings.ContainsRune,
				"Count":          strings.Count,
				"EqualFold":      strings.EqualFold,
				"Fields":         strings.Fields,
				"FieldsFunc":     strings.FieldsFunc,
				"HasPrefix":      strings.HasPrefix,
				"HasSuffix":      strings.HasSuffix,
				"Index":          strings.Index,
				"IndexAny":       strings.IndexAny,
				"IndexByte":      strings.IndexByte,
				"IndexFunc":      strings.IndexFunc,
				"IndexRune":      strings.IndexRune,
				"Join":           strings.Join,
				"LastIndex":      strings.LastIndex,
				"LastIndexAny":   strings.LastIndexAny,
				"LastIndexByte":  strings.LastIndexByte,
				"LastIndexFunc":  strings.LastIndexFunc,
				"Map":            strings.Map,
				"NewReader":      strings.NewReader,
				"NewReplacer":    strings.NewReplacer,
				"Reader":         reflect.TypeOf(strings.Reader{}),
				"Repeat":         strings.Repeat,
				"Replace":        strings.Replace,
				"ReplaceAll":     strings.ReplaceAll,
				"Replacer":       reflect.TypeOf(strings.Replacer{}),
				"Split":          strings.Split,
				"SplitAfter":     strings.SplitAfter,
				"SplitAfterN":    strings.SplitAfterN,
				"SplitN":         strings.SplitN,
				"Title":          strings.Title,
				"ToLower":        strings.ToLower,
				"ToLowerSpecial": strings.ToLowerSpecial,
				"ToTitle":        strings.ToTitle,
				"ToTitleSpecial": strings.ToTitleSpecial,
				"ToUpper":        strings.ToUpper,
				"ToUpperSpecial": strings.ToUpperSpecial,
				"Trim":           strings.Trim,
				"TrimFunc":       strings.TrimFunc,
				"TrimLeft":       strings.TrimLeft,
				"TrimLeftFunc":   strings.TrimLeftFunc,
				"TrimPrefix":     strings.TrimPrefix,
				"TrimRight":      strings.TrimRight,
				"TrimRightFunc":  strings.TrimRightFunc,
				"TrimSpace":      strings.TrimSpace,
				"TrimSuffix":     strings.TrimSuffix,
			},
		},
		"time": {
			Name: "time",
			Declarations: map[string]interface{}{
				"ANSIC":                  ConstValue(time.ANSIC),
				"After":                  time.After,
				"AfterFunc":              time.AfterFunc,
				"April":                  ConstValue(time.April),
				"August":                 ConstValue(time.August),
				"Date":                   time.Date,
				"December":               ConstValue(time.December),
				"Duration":               reflect.TypeOf(new(time.Duration)).Elem(),
				"February":               ConstValue(time.February),
				"FixedZone":              time.FixedZone,
				"Friday":                 ConstValue(time.Friday),
				"Hour":                   ConstLiteral(reflect.TypeOf(new(time.Duration)).Elem(), "3600000000000"),
				"January":                ConstValue(time.January),
				"July":                   ConstValue(time.July),
				"June":                   ConstValue(time.June),
				"Kitchen":                ConstValue(time.Kitchen),
				"LoadLocation":           time.LoadLocation,
				"LoadLocationFromTZData": time.LoadLocationFromTZData,
				"Local":                  &time.Local,
				"Location":               reflect.TypeOf(time.Location{}),
				"March":                  ConstValue(time.March),
				"May":                    ConstValue(time.May),
				"Microsecond":            ConstValue(time.Microsecond),
				"Millisecond":            ConstLiteral(reflect.TypeOf(new(time.Duration)).Elem(), "1000000"),
				"Minute":                 ConstLiteral(reflect.TypeOf(new(time.Duration)).Elem(), "60000000000"),
				"Monday":                 ConstValue(time.Monday),
				"Month":                  reflect.TypeOf(new(time.Month)).Elem(),
				"Nanosecond":             ConstValue(time.Nanosecond),
				"NewTicker":              time.NewTicker,
				"NewTimer":               time.NewTimer,
				"November":               ConstValue(time.November),
				"Now":                    time.Now,
				"October":                ConstValue(time.October),
				"Parse":                  time.Parse,
				"ParseDuration":          time.ParseDuration,
				"ParseError":             reflect.TypeOf(time.ParseError{}),
				"ParseInLocation":        time.ParseInLocation,
				"RFC1123":                ConstValue(time.RFC1123),
				"RFC1123Z":               ConstValue(time.RFC1123Z),
				"RFC3339":                ConstValue(time.RFC3339),
				"RFC3339Nano":            ConstValue(time.RFC3339Nano),
				"RFC822":                 ConstValue(time.RFC822),
				"RFC822Z":                ConstValue(time.RFC822Z),
				"RFC850":                 ConstValue(time.RFC850),
				"RubyDate":               ConstValue(time.RubyDate),
				"Saturday":               ConstValue(time.Saturday),
				"Second":                 ConstLiteral(reflect.TypeOf(new(time.Duration)).Elem(), "1000000000"),
				"September":              ConstValue(time.September),
				"Since":                  time.Since,
				"Sleep":                  time.Sleep,
				"Stamp":                  ConstValue(time.Stamp),
				"StampMicro":             ConstValue(time.StampMicro),
				"StampMilli":             ConstValue(time.StampMilli),
				"StampNano":              ConstValue(time.StampNano),
				"Sunday":                 ConstValue(time.Sunday),
				"Thursday":               ConstValue(time.Thursday),
				"Tick":                   time.Tick,
				"Ticker":                 reflect.TypeOf(time.Ticker{}),
				"Time":                   reflect.TypeOf(time.Time{}),
				"Timer":                  reflect.TypeOf(time.Timer{}),
				"Tuesday":                ConstValue(time.Tuesday),
				"UTC":                    &time.UTC,
				"Unix":                   time.Unix,
				"UnixDate":               ConstValue(time.UnixDate),
				"Until":                  time.Until,
				"Wednesday":              ConstValue(time.Wednesday),
				"Weekday":                reflect.TypeOf(new(time.Weekday)).Elem(),
			},
		},
		"bytes": {
			Name: "bytes",
			Declarations: map[string]interface{}{
				"Buffer":          reflect.TypeOf(bytes.Buffer{}),
				"Compare":         bytes.Compare,
				"Contains":        bytes.Contains,
				"ContainsAny":     bytes.ContainsAny,
				"ContainsRune":    bytes.ContainsRune,
				"Count":           bytes.Count,
				"Equal":           bytes.Equal,
				"EqualFold":       bytes.EqualFold,
				"ErrTooLarge":     &bytes.ErrTooLarge,
				"Fields":          bytes.Fields,
				"FieldsFunc":      bytes.FieldsFunc,
				"HasPrefix":       bytes.HasPrefix,
				"HasSuffix":       bytes.HasSuffix,
				"Index":           bytes.Index,
				"IndexAny":        bytes.IndexAny,
				"IndexByte":       bytes.IndexByte,
				"IndexFunc":       bytes.IndexFunc,
				"IndexRune":       bytes.IndexRune,
				"Join":            bytes.Join,
				"LastIndex":       bytes.LastIndex,
				"LastIndexAny":    bytes.LastIndexAny,
				"LastIndexByte":   bytes.LastIndexByte,
				"LastIndexFunc":   bytes.LastIndexFunc,
				"Map":             bytes.Map,
				"MinRead":         ConstValue(bytes.MinRead),
				"NewBuffer":       bytes.NewBuffer,
				"NewBufferString": bytes.NewBufferString,
				"NewReader":       bytes.NewReader,
				"Reader":          reflect.TypeOf(bytes.Reader{}),
				"Repeat":          bytes.Repeat,
				"Replace":         bytes.Replace,
				"ReplaceAll":      bytes.ReplaceAll,
				"Runes":           bytes.Runes,
				"Split":           bytes.Split,
				"SplitAfter":      bytes.SplitAfter,
				"SplitAfterN":     bytes.SplitAfterN,
				"SplitN":          bytes.SplitN,
				"Title":           bytes.Title,
				"ToLower":         bytes.ToLower,
				"ToLowerSpecial":  bytes.ToLowerSpecial,
				"ToTitle":         bytes.ToTitle,
				"ToTitleSpecial":  bytes.ToTitleSpecial,
				"ToUpper":         bytes.ToUpper,
				"ToUpperSpecial":  bytes.ToUpperSpecial,
				"Trim":            bytes.Trim,
				"TrimFunc":        bytes.TrimFunc,
				"TrimLeft":        bytes.TrimLeft,
				"TrimLeftFunc":    bytes.TrimLeftFunc,
				"TrimPrefix":      bytes.TrimPrefix,
				"TrimRight":       bytes.TrimRight,
				"TrimRightFunc":   bytes.TrimRightFunc,
				"TrimSpace":       bytes.TrimSpace,
				"TrimSuffix":      bytes.TrimSuffix,
			},
		},
		"crypto/sha1": {
			Name: "sha1",
			Declarations: map[string]interface{}{
				"BlockSize": ConstValue(sha1.BlockSize),
				"New":       sha1.New,
				"Size":      ConstValue(sha1.Size),
				"Sum":       sha1.Sum,
			},
		},
	}
}
