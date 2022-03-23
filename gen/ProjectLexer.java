// Generated from /home/jdvelasquezo/GolandProjects/OLC2_Project1/server/ProjectLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProjectLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, PRINT=2, DECLARAR=3, MUT=4, RSTRING=5, RINTEGER=6, RREAL=7, 
		RBOOLEAN=8, RCHAR=9, RSTR=10, RAS=11, RIF=12, RELSE=13, RMATCH=14, RWHILE=15, 
		RLOOP=16, RFOR=17, RIN=18, RBREAK=19, RCONTINUE=20, RRETURN=21, POWI=22, 
		POWF=23, TOSTRING=24, TOOWNED=25, RMAIN=26, RFN=27, RVECTOR=28, RVEC=29, 
		RVECMayus=30, RNEW=31, REVECTORNEW=32, REVECCAPACITY=33, RSTRUCT=34, MULTILINE=35, 
		INLINE=36, INTEGER=37, FLOAT=38, CHAR=39, STRING=40, BOOLEAN=41, ID=42, 
		EQUAL=43, DOT=44, SEMICOLON=45, COMMA=46, COLON=47, ADMIRATION=48, REFERENCE=49, 
		HERITAGE=50, ARROW=51, RANGE=52, LEFT_PAR=53, RIGHT_PAR=54, LEFT_KEY=55, 
		RIGHT_KEY=56, LEFT_BRACKET=57, RIGHT_BRACKET=58, EQUEAL_EQUAL=59, NOTEQUAL=60, 
		GREATER_THAN=61, LESS_THAN=62, GREATER_EQUALTHAN=63, LESS_EQUEALTHAN=64, 
		MUL=65, DIV=66, MOD=67, ADD=68, SUB=69, AND=70, OR=71, PIPE=72, EQUAL_ARROW=73, 
		UNDERSCORE=74, WHITESPACE=75;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", 
			"TOSTRING", "TOOWNED", "RMAIN", "RFN", "RVECTOR", "RVEC", "RVECMayus", 
			"RNEW", "REVECTORNEW", "REVECCAPACITY", "RSTRUCT", "MULTILINE", "INLINE", 
			"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", 
			"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", 
			"ARROW", "RANGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", 
			"RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", 
			"GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", 
			"AND", "OR", "PIPE", "EQUAL_ARROW", "UNDERSCORE", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", 
			"'f64'", "'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'match'", 
			"'while'", "'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'", 
			"'pow'", "'powf'", "'to_string'", "'to_owned'", "'main'", "'fn'", "'vector'", 
			"'vec'", "'Vec'", "'new'", "'Vec::new()'", "'Vec::with_capacity'", "'struct'", 
			null, null, null, null, null, null, null, null, "'='", "'.'", "';'", 
			"','", "':'", "'!'", "'&'", "'::'", "'->'", "'..'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'*'", 
			"'/'", "'%'", "'+'", "'-'", "'&&'", "'||'", "'|'", "'=>'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", 
			"TOSTRING", "TOOWNED", "RMAIN", "RFN", "RVECTOR", "RVEC", "RVECMayus", 
			"RNEW", "REVECTORNEW", "REVECCAPACITY", "RSTRUCT", "MULTILINE", "INLINE", 
			"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", 
			"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", 
			"ARROW", "RANGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", 
			"RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", 
			"GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", 
			"AND", "OR", "PIPE", "EQUAL_ARROW", "UNDERSCORE", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ProjectLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ProjectLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2M\u01fe\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30"+
		"\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37"+
		"\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3"+
		"#\3#\3#\3$\3$\3$\3$\6$\u0169\n$\r$\16$\u016a\3$\3$\3$\3$\3$\3%\3%\3%\3"+
		"%\6%\u0176\n%\r%\16%\u0177\3%\3%\3&\6&\u017d\n&\r&\16&\u017e\3\'\6\'\u0182"+
		"\n\'\r\'\16\'\u0183\3\'\3\'\6\'\u0188\n\'\r\'\16\'\u0189\3(\3(\3(\3(\3"+
		")\3)\7)\u0192\n)\f)\16)\u0195\13)\3)\3)\3*\3*\3*\3*\3*\3*\3*\3*\3*\5*"+
		"\u01a2\n*\3+\3+\7+\u01a6\n+\f+\16+\u01a9\13+\3,\3,\3-\3-\3.\3.\3/\3/\3"+
		"\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\64\3\65\3\65\3"+
		"\65\3\66\3\66\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3<\3=\3=\3=\3>\3"+
		">\3?\3?\3@\3@\3@\3A\3A\3A\3B\3B\3C\3C\3D\3D\3E\3E\3F\3F\3G\3G\3G\3H\3"+
		"H\3H\3I\3I\3J\3J\3J\3K\3K\3L\6L\u01f6\nL\rL\16L\u01f7\3L\3L\3M\3M\3M\2"+
		"\2N\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67"+
		"m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008d"+
		"H\u008fI\u0091J\u0093K\u0095L\u0097M\u0099\2\3\2\n\3\2\61\61\3\2\f\f\3"+
		"\2\62;\3\2$$\7\2C\\aac|\u00d3\u00d3\u00f3\u00f3\b\2\62;C\\aac|\u00d3\u00d3"+
		"\u00f3\u00f3\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u0205\2\3\3\2\2"+
		"\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3"+
		"\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2"+
		"\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2"+
		"\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2"+
		"\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3"+
		"\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2"+
		"\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2"+
		"W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3"+
		"\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2"+
		"\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2"+
		"}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2"+
		"\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f"+
		"\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2\2\2\u0097\3\2\2"+
		"\2\3\u009b\3\2\2\2\5\u00a3\3\2\2\2\7\u00a9\3\2\2\2\t\u00ad\3\2\2\2\13"+
		"\u00b1\3\2\2\2\r\u00b8\3\2\2\2\17\u00bc\3\2\2\2\21\u00c0\3\2\2\2\23\u00c5"+
		"\3\2\2\2\25\u00ca\3\2\2\2\27\u00cf\3\2\2\2\31\u00d2\3\2\2\2\33\u00d5\3"+
		"\2\2\2\35\u00da\3\2\2\2\37\u00e0\3\2\2\2!\u00e6\3\2\2\2#\u00eb\3\2\2\2"+
		"%\u00ef\3\2\2\2\'\u00f2\3\2\2\2)\u00f8\3\2\2\2+\u0101\3\2\2\2-\u0108\3"+
		"\2\2\2/\u010c\3\2\2\2\61\u0111\3\2\2\2\63\u011b\3\2\2\2\65\u0124\3\2\2"+
		"\2\67\u0129\3\2\2\29\u012c\3\2\2\2;\u0133\3\2\2\2=\u0137\3\2\2\2?\u013b"+
		"\3\2\2\2A\u013f\3\2\2\2C\u014a\3\2\2\2E\u015d\3\2\2\2G\u0164\3\2\2\2I"+
		"\u0171\3\2\2\2K\u017c\3\2\2\2M\u0181\3\2\2\2O\u018b\3\2\2\2Q\u018f\3\2"+
		"\2\2S\u01a1\3\2\2\2U\u01a3\3\2\2\2W\u01aa\3\2\2\2Y\u01ac\3\2\2\2[\u01ae"+
		"\3\2\2\2]\u01b0\3\2\2\2_\u01b2\3\2\2\2a\u01b4\3\2\2\2c\u01b6\3\2\2\2e"+
		"\u01b8\3\2\2\2g\u01bb\3\2\2\2i\u01be\3\2\2\2k\u01c1\3\2\2\2m\u01c3\3\2"+
		"\2\2o\u01c5\3\2\2\2q\u01c7\3\2\2\2s\u01c9\3\2\2\2u\u01cb\3\2\2\2w\u01cd"+
		"\3\2\2\2y\u01d0\3\2\2\2{\u01d3\3\2\2\2}\u01d5\3\2\2\2\177\u01d7\3\2\2"+
		"\2\u0081\u01da\3\2\2\2\u0083\u01dd\3\2\2\2\u0085\u01df\3\2\2\2\u0087\u01e1"+
		"\3\2\2\2\u0089\u01e3\3\2\2\2\u008b\u01e5\3\2\2\2\u008d\u01e7\3\2\2\2\u008f"+
		"\u01ea\3\2\2\2\u0091\u01ed\3\2\2\2\u0093\u01ef\3\2\2\2\u0095\u01f2\3\2"+
		"\2\2\u0097\u01f5\3\2\2\2\u0099\u01fb\3\2\2\2\u009b\u009c\7r\2\2\u009c"+
		"\u009d\7t\2\2\u009d\u009e\7k\2\2\u009e\u009f\7p\2\2\u009f\u00a0\7v\2\2"+
		"\u00a0\u00a1\7n\2\2\u00a1\u00a2\7p\2\2\u00a2\4\3\2\2\2\u00a3\u00a4\7r"+
		"\2\2\u00a4\u00a5\7t\2\2\u00a5\u00a6\7k\2\2\u00a6\u00a7\7p\2\2\u00a7\u00a8"+
		"\7v\2\2\u00a8\6\3\2\2\2\u00a9\u00aa\7n\2\2\u00aa\u00ab\7g\2\2\u00ab\u00ac"+
		"\7v\2\2\u00ac\b\3\2\2\2\u00ad\u00ae\7o\2\2\u00ae\u00af\7w\2\2\u00af\u00b0"+
		"\7v\2\2\u00b0\n\3\2\2\2\u00b1\u00b2\7U\2\2\u00b2\u00b3\7v\2\2\u00b3\u00b4"+
		"\7t\2\2\u00b4\u00b5\7k\2\2\u00b5\u00b6\7p\2\2\u00b6\u00b7\7i\2\2\u00b7"+
		"\f\3\2\2\2\u00b8\u00b9\7k\2\2\u00b9\u00ba\78\2\2\u00ba\u00bb\7\66\2\2"+
		"\u00bb\16\3\2\2\2\u00bc\u00bd\7h\2\2\u00bd\u00be\78\2\2\u00be\u00bf\7"+
		"\66\2\2\u00bf\20\3\2\2\2\u00c0\u00c1\7d\2\2\u00c1\u00c2\7q\2\2\u00c2\u00c3"+
		"\7q\2\2\u00c3\u00c4\7n\2\2\u00c4\22\3\2\2\2\u00c5\u00c6\7e\2\2\u00c6\u00c7"+
		"\7j\2\2\u00c7\u00c8\7c\2\2\u00c8\u00c9\7t\2\2\u00c9\24\3\2\2\2\u00ca\u00cb"+
		"\7(\2\2\u00cb\u00cc\7u\2\2\u00cc\u00cd\7v\2\2\u00cd\u00ce\7t\2\2\u00ce"+
		"\26\3\2\2\2\u00cf\u00d0\7c\2\2\u00d0\u00d1\7u\2\2\u00d1\30\3\2\2\2\u00d2"+
		"\u00d3\7k\2\2\u00d3\u00d4\7h\2\2\u00d4\32\3\2\2\2\u00d5\u00d6\7g\2\2\u00d6"+
		"\u00d7\7n\2\2\u00d7\u00d8\7u\2\2\u00d8\u00d9\7g\2\2\u00d9\34\3\2\2\2\u00da"+
		"\u00db\7o\2\2\u00db\u00dc\7c\2\2\u00dc\u00dd\7v\2\2\u00dd\u00de\7e\2\2"+
		"\u00de\u00df\7j\2\2\u00df\36\3\2\2\2\u00e0\u00e1\7y\2\2\u00e1\u00e2\7"+
		"j\2\2\u00e2\u00e3\7k\2\2\u00e3\u00e4\7n\2\2\u00e4\u00e5\7g\2\2\u00e5 "+
		"\3\2\2\2\u00e6\u00e7\7n\2\2\u00e7\u00e8\7q\2\2\u00e8\u00e9\7q\2\2\u00e9"+
		"\u00ea\7r\2\2\u00ea\"\3\2\2\2\u00eb\u00ec\7h\2\2\u00ec\u00ed\7q\2\2\u00ed"+
		"\u00ee\7t\2\2\u00ee$\3\2\2\2\u00ef\u00f0\7k\2\2\u00f0\u00f1\7p\2\2\u00f1"+
		"&\3\2\2\2\u00f2\u00f3\7d\2\2\u00f3\u00f4\7t\2\2\u00f4\u00f5\7g\2\2\u00f5"+
		"\u00f6\7c\2\2\u00f6\u00f7\7m\2\2\u00f7(\3\2\2\2\u00f8\u00f9\7e\2\2\u00f9"+
		"\u00fa\7q\2\2\u00fa\u00fb\7p\2\2\u00fb\u00fc\7v\2\2\u00fc\u00fd\7k\2\2"+
		"\u00fd\u00fe\7p\2\2\u00fe\u00ff\7w\2\2\u00ff\u0100\7g\2\2\u0100*\3\2\2"+
		"\2\u0101\u0102\7t\2\2\u0102\u0103\7g\2\2\u0103\u0104\7v\2\2\u0104\u0105"+
		"\7w\2\2\u0105\u0106\7t\2\2\u0106\u0107\7p\2\2\u0107,\3\2\2\2\u0108\u0109"+
		"\7r\2\2\u0109\u010a\7q\2\2\u010a\u010b\7y\2\2\u010b.\3\2\2\2\u010c\u010d"+
		"\7r\2\2\u010d\u010e\7q\2\2\u010e\u010f\7y\2\2\u010f\u0110\7h\2\2\u0110"+
		"\60\3\2\2\2\u0111\u0112\7v\2\2\u0112\u0113\7q\2\2\u0113\u0114\7a\2\2\u0114"+
		"\u0115\7u\2\2\u0115\u0116\7v\2\2\u0116\u0117\7t\2\2\u0117\u0118\7k\2\2"+
		"\u0118\u0119\7p\2\2\u0119\u011a\7i\2\2\u011a\62\3\2\2\2\u011b\u011c\7"+
		"v\2\2\u011c\u011d\7q\2\2\u011d\u011e\7a\2\2\u011e\u011f\7q\2\2\u011f\u0120"+
		"\7y\2\2\u0120\u0121\7p\2\2\u0121\u0122\7g\2\2\u0122\u0123\7f\2\2\u0123"+
		"\64\3\2\2\2\u0124\u0125\7o\2\2\u0125\u0126\7c\2\2\u0126\u0127\7k\2\2\u0127"+
		"\u0128\7p\2\2\u0128\66\3\2\2\2\u0129\u012a\7h\2\2\u012a\u012b\7p\2\2\u012b"+
		"8\3\2\2\2\u012c\u012d\7x\2\2\u012d\u012e\7g\2\2\u012e\u012f\7e\2\2\u012f"+
		"\u0130\7v\2\2\u0130\u0131\7q\2\2\u0131\u0132\7t\2\2\u0132:\3\2\2\2\u0133"+
		"\u0134\7x\2\2\u0134\u0135\7g\2\2\u0135\u0136\7e\2\2\u0136<\3\2\2\2\u0137"+
		"\u0138\7X\2\2\u0138\u0139\7g\2\2\u0139\u013a\7e\2\2\u013a>\3\2\2\2\u013b"+
		"\u013c\7p\2\2\u013c\u013d\7g\2\2\u013d\u013e\7y\2\2\u013e@\3\2\2\2\u013f"+
		"\u0140\7X\2\2\u0140\u0141\7g\2\2\u0141\u0142\7e\2\2\u0142\u0143\7<\2\2"+
		"\u0143\u0144\7<\2\2\u0144\u0145\7p\2\2\u0145\u0146\7g\2\2\u0146\u0147"+
		"\7y\2\2\u0147\u0148\7*\2\2\u0148\u0149\7+\2\2\u0149B\3\2\2\2\u014a\u014b"+
		"\7X\2\2\u014b\u014c\7g\2\2\u014c\u014d\7e\2\2\u014d\u014e\7<\2\2\u014e"+
		"\u014f\7<\2\2\u014f\u0150\7y\2\2\u0150\u0151\7k\2\2\u0151\u0152\7v\2\2"+
		"\u0152\u0153\7j\2\2\u0153\u0154\7a\2\2\u0154\u0155\7e\2\2\u0155\u0156"+
		"\7c\2\2\u0156\u0157\7r\2\2\u0157\u0158\7c\2\2\u0158\u0159\7e\2\2\u0159"+
		"\u015a\7k\2\2\u015a\u015b\7v\2\2\u015b\u015c\7{\2\2\u015cD\3\2\2\2\u015d"+
		"\u015e\7u\2\2\u015e\u015f\7v\2\2\u015f\u0160\7t\2\2\u0160\u0161\7w\2\2"+
		"\u0161\u0162\7e\2\2\u0162\u0163\7v\2\2\u0163F\3\2\2\2\u0164\u0165\7\61"+
		"\2\2\u0165\u0166\7,\2\2\u0166\u0168\3\2\2\2\u0167\u0169\n\2\2\2\u0168"+
		"\u0167\3\2\2\2\u0169\u016a\3\2\2\2\u016a\u0168\3\2\2\2\u016a\u016b\3\2"+
		"\2\2\u016b\u016c\3\2\2\2\u016c\u016d\7,\2\2\u016d\u016e\7\61\2\2\u016e"+
		"\u016f\3\2\2\2\u016f\u0170\b$\2\2\u0170H\3\2\2\2\u0171\u0172\7\61\2\2"+
		"\u0172\u0173\7\61\2\2\u0173\u0175\3\2\2\2\u0174\u0176\n\3\2\2\u0175\u0174"+
		"\3\2\2\2\u0176\u0177\3\2\2\2\u0177\u0175\3\2\2\2\u0177\u0178\3\2\2\2\u0178"+
		"\u0179\3\2\2\2\u0179\u017a\b%\2\2\u017aJ\3\2\2\2\u017b\u017d\t\4\2\2\u017c"+
		"\u017b\3\2\2\2\u017d\u017e\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2"+
		"\2\2\u017fL\3\2\2\2\u0180\u0182\t\4\2\2\u0181\u0180\3\2\2\2\u0182\u0183"+
		"\3\2\2\2\u0183\u0181\3\2\2\2\u0183\u0184\3\2\2\2\u0184\u0185\3\2\2\2\u0185"+
		"\u0187\7\60\2\2\u0186\u0188\t\4\2\2\u0187\u0186\3\2\2\2\u0188\u0189\3"+
		"\2\2\2\u0189\u0187\3\2\2\2\u0189\u018a\3\2\2\2\u018aN\3\2\2\2\u018b\u018c"+
		"\7)\2\2\u018c\u018d\n\5\2\2\u018d\u018e\7)\2\2\u018eP\3\2\2\2\u018f\u0193"+
		"\7$\2\2\u0190\u0192\n\5\2\2\u0191\u0190\3\2\2\2\u0192\u0195\3\2\2\2\u0193"+
		"\u0191\3\2\2\2\u0193\u0194\3\2\2\2\u0194\u0196\3\2\2\2\u0195\u0193\3\2"+
		"\2\2\u0196\u0197\7$\2\2\u0197R\3\2\2\2\u0198\u0199\7v\2\2\u0199\u019a"+
		"\7t\2\2\u019a\u019b\7w\2\2\u019b\u01a2\7g\2\2\u019c\u019d\7h\2\2\u019d"+
		"\u019e\7c\2\2\u019e\u019f\7n\2\2\u019f\u01a0\7u\2\2\u01a0\u01a2\7g\2\2"+
		"\u01a1\u0198\3\2\2\2\u01a1\u019c\3\2\2\2\u01a2T\3\2\2\2\u01a3\u01a7\t"+
		"\6\2\2\u01a4\u01a6\t\7\2\2\u01a5\u01a4\3\2\2\2\u01a6\u01a9\3\2\2\2\u01a7"+
		"\u01a5\3\2\2\2\u01a7\u01a8\3\2\2\2\u01a8V\3\2\2\2\u01a9\u01a7\3\2\2\2"+
		"\u01aa\u01ab\7?\2\2\u01abX\3\2\2\2\u01ac\u01ad\7\60\2\2\u01adZ\3\2\2\2"+
		"\u01ae\u01af\7=\2\2\u01af\\\3\2\2\2\u01b0\u01b1\7.\2\2\u01b1^\3\2\2\2"+
		"\u01b2\u01b3\7<\2\2\u01b3`\3\2\2\2\u01b4\u01b5\7#\2\2\u01b5b\3\2\2\2\u01b6"+
		"\u01b7\7(\2\2\u01b7d\3\2\2\2\u01b8\u01b9\7<\2\2\u01b9\u01ba\7<\2\2\u01ba"+
		"f\3\2\2\2\u01bb\u01bc\7/\2\2\u01bc\u01bd\7@\2\2\u01bdh\3\2\2\2\u01be\u01bf"+
		"\7\60\2\2\u01bf\u01c0\7\60\2\2\u01c0j\3\2\2\2\u01c1\u01c2\7*\2\2\u01c2"+
		"l\3\2\2\2\u01c3\u01c4\7+\2\2\u01c4n\3\2\2\2\u01c5\u01c6\7}\2\2\u01c6p"+
		"\3\2\2\2\u01c7\u01c8\7\177\2\2\u01c8r\3\2\2\2\u01c9\u01ca\7]\2\2\u01ca"+
		"t\3\2\2\2\u01cb\u01cc\7_\2\2\u01ccv\3\2\2\2\u01cd\u01ce\7?\2\2\u01ce\u01cf"+
		"\7?\2\2\u01cfx\3\2\2\2\u01d0\u01d1\7#\2\2\u01d1\u01d2\7?\2\2\u01d2z\3"+
		"\2\2\2\u01d3\u01d4\7@\2\2\u01d4|\3\2\2\2\u01d5\u01d6\7>\2\2\u01d6~\3\2"+
		"\2\2\u01d7\u01d8\7@\2\2\u01d8\u01d9\7?\2\2\u01d9\u0080\3\2\2\2\u01da\u01db"+
		"\7>\2\2\u01db\u01dc\7?\2\2\u01dc\u0082\3\2\2\2\u01dd\u01de\7,\2\2\u01de"+
		"\u0084\3\2\2\2\u01df\u01e0\7\61\2\2\u01e0\u0086\3\2\2\2\u01e1\u01e2\7"+
		"\'\2\2\u01e2\u0088\3\2\2\2\u01e3\u01e4\7-\2\2\u01e4\u008a\3\2\2\2\u01e5"+
		"\u01e6\7/\2\2\u01e6\u008c\3\2\2\2\u01e7\u01e8\7(\2\2\u01e8\u01e9\7(\2"+
		"\2\u01e9\u008e\3\2\2\2\u01ea\u01eb\7~\2\2\u01eb\u01ec\7~\2\2\u01ec\u0090"+
		"\3\2\2\2\u01ed\u01ee\7~\2\2\u01ee\u0092\3\2\2\2\u01ef\u01f0\7?\2\2\u01f0"+
		"\u01f1\7@\2\2\u01f1\u0094\3\2\2\2\u01f2\u01f3\7a\2\2\u01f3\u0096\3\2\2"+
		"\2\u01f4\u01f6\t\b\2\2\u01f5\u01f4\3\2\2\2\u01f6\u01f7\3\2\2\2\u01f7\u01f5"+
		"\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8\u01f9\3\2\2\2\u01f9\u01fa\bL\2\2\u01fa"+
		"\u0098\3\2\2\2\u01fb\u01fc\7^\2\2\u01fc\u01fd\t\t\2\2\u01fd\u009a\3\2"+
		"\2\2\f\2\u016a\u0177\u017e\u0183\u0189\u0193\u01a1\u01a7\u01f7\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}