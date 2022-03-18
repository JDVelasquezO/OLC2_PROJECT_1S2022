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
		RBREAK=16, RCONTINUE=17, RRETURN=18, POWI=19, POWF=20, TOSTRING=21, TOOWNED=22, 
		RMAIN=23, RFN=24, MULTILINE=25, INLINE=26, INTEGER=27, FLOAT=28, CHAR=29, 
		STRING=30, BOOLEAN=31, ID=32, EQUAL=33, DOT=34, SEMICOLON=35, COMMA=36, 
		COLON=37, ADMIRATION=38, REFERENCE=39, HERITAGE=40, ARROW=41, LEFT_PAR=42, 
		RIGHT_PAR=43, LEFT_KEY=44, RIGHT_KEY=45, LEFT_BRACKET=46, RIGHT_BRACKET=47, 
		EQUEAL_EQUAL=48, NOTEQUAL=49, GREATER_THAN=50, LESS_THAN=51, GREATER_EQUALTHAN=52, 
		LESS_EQUEALTHAN=53, MUL=54, DIV=55, MOD=56, ADD=57, SUB=58, AND=59, OR=60, 
		PIPE=61, EQUAL_ARROW=62, UNDERSCORE=63, WHITESPACE=64;
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
			"RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", "TOSTRING", "TOOWNED", 
			"RMAIN", "RFN", "MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", 
			"BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", 
			"REFERENCE", "HERITAGE", "ARROW", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", 
			"RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", 
			"GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", 
			"MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "PIPE", "EQUAL_ARROW", 
			"UNDERSCORE", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", 
			"'f64'", "'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'match'", 
			"'while'", "'break'", "'continue'", "'return'", "'pow'", "'powf'", "'to_string'", 
			"'to_owned'", "'main'", "'fn'", null, null, null, null, null, null, null, 
			null, "'='", "'.'", "';'", "','", "':'", "'!'", "'&'", "'::'", "'->'", 
			"'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", "'>'", "'<'", 
			"'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'", "'&&'", "'||'", "'|'", 
			"'=>'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", "TOSTRING", "TOOWNED", 
			"RMAIN", "RFN", "MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING", 
			"BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", 
			"REFERENCE", "HERITAGE", "ARROW", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", 
			"RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", 
			"GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", 
			"MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "PIPE", "EQUAL_ARROW", 
			"UNDERSCORE", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2B\u01a1\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n"+
		"\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3"+
		"\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3"+
		"\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3"+
		"\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3"+
		"\31\3\31\3\31\3\32\3\32\3\32\3\32\6\32\u010f\n\32\r\32\16\32\u0110\3\32"+
		"\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\6\33\u011c\n\33\r\33\16\33\u011d"+
		"\3\33\3\33\3\34\6\34\u0123\n\34\r\34\16\34\u0124\3\35\6\35\u0128\n\35"+
		"\r\35\16\35\u0129\3\35\3\35\6\35\u012e\n\35\r\35\16\35\u012f\3\36\3\36"+
		"\3\36\3\36\3\37\3\37\7\37\u0138\n\37\f\37\16\37\u013b\13\37\3\37\3\37"+
		"\3 \3 \3 \3 \3 \3 \3 \3 \3 \5 \u0148\n \3!\3!\7!\u014c\n!\f!\16!\u014f"+
		"\13!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3)\3*\3*\3*\3"+
		"+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\61\3\62\3\62\3\62\3"+
		"\63\3\63\3\64\3\64\3\65\3\65\3\65\3\66\3\66\3\66\3\67\3\67\38\38\39\3"+
		"9\3:\3:\3;\3;\3<\3<\3<\3=\3=\3=\3>\3>\3?\3?\3?\3@\3@\3A\6A\u0199\nA\r"+
		"A\16A\u019a\3A\3A\3B\3B\3B\2\2C\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31"+
		"\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60"+
		"_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083\2\3\2"+
		"\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7\2C\\aac|\u00d3\u00d3\u00f3\u00f3"+
		"\b\2\62;C\\aac|\u00d3\u00d3\u00f3\u00f3\5\2\13\f\17\17\"\"\t\2\"#%%--"+
		"/\60<<BB]_\2\u01a8\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13"+
		"\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2"+
		"\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2"+
		"!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3"+
		"\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2"+
		"\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E"+
		"\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2"+
		"\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2"+
		"\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k"+
		"\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2"+
		"\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\3"+
		"\u0085\3\2\2\2\5\u008d\3\2\2\2\7\u0093\3\2\2\2\t\u0097\3\2\2\2\13\u009b"+
		"\3\2\2\2\r\u00a2\3\2\2\2\17\u00a6\3\2\2\2\21\u00aa\3\2\2\2\23\u00af\3"+
		"\2\2\2\25\u00b4\3\2\2\2\27\u00b9\3\2\2\2\31\u00bc\3\2\2\2\33\u00bf\3\2"+
		"\2\2\35\u00c4\3\2\2\2\37\u00ca\3\2\2\2!\u00d0\3\2\2\2#\u00d6\3\2\2\2%"+
		"\u00df\3\2\2\2\'\u00e6\3\2\2\2)\u00ea\3\2\2\2+\u00ef\3\2\2\2-\u00f9\3"+
		"\2\2\2/\u0102\3\2\2\2\61\u0107\3\2\2\2\63\u010a\3\2\2\2\65\u0117\3\2\2"+
		"\2\67\u0122\3\2\2\29\u0127\3\2\2\2;\u0131\3\2\2\2=\u0135\3\2\2\2?\u0147"+
		"\3\2\2\2A\u0149\3\2\2\2C\u0150\3\2\2\2E\u0152\3\2\2\2G\u0154\3\2\2\2I"+
		"\u0156\3\2\2\2K\u0158\3\2\2\2M\u015a\3\2\2\2O\u015c\3\2\2\2Q\u015e\3\2"+
		"\2\2S\u0161\3\2\2\2U\u0164\3\2\2\2W\u0166\3\2\2\2Y\u0168\3\2\2\2[\u016a"+
		"\3\2\2\2]\u016c\3\2\2\2_\u016e\3\2\2\2a\u0170\3\2\2\2c\u0173\3\2\2\2e"+
		"\u0176\3\2\2\2g\u0178\3\2\2\2i\u017a\3\2\2\2k\u017d\3\2\2\2m\u0180\3\2"+
		"\2\2o\u0182\3\2\2\2q\u0184\3\2\2\2s\u0186\3\2\2\2u\u0188\3\2\2\2w\u018a"+
		"\3\2\2\2y\u018d\3\2\2\2{\u0190\3\2\2\2}\u0192\3\2\2\2\177\u0195\3\2\2"+
		"\2\u0081\u0198\3\2\2\2\u0083\u019e\3\2\2\2\u0085\u0086\7r\2\2\u0086\u0087"+
		"\7t\2\2\u0087\u0088\7k\2\2\u0088\u0089\7p\2\2\u0089\u008a\7v\2\2\u008a"+
		"\u008b\7n\2\2\u008b\u008c\7p\2\2\u008c\4\3\2\2\2\u008d\u008e\7r\2\2\u008e"+
		"\u008f\7t\2\2\u008f\u0090\7k\2\2\u0090\u0091\7p\2\2\u0091\u0092\7v\2\2"+
		"\u0092\6\3\2\2\2\u0093\u0094\7n\2\2\u0094\u0095\7g\2\2\u0095\u0096\7v"+
		"\2\2\u0096\b\3\2\2\2\u0097\u0098\7o\2\2\u0098\u0099\7w\2\2\u0099\u009a"+
		"\7v\2\2\u009a\n\3\2\2\2\u009b\u009c\7U\2\2\u009c\u009d\7v\2\2\u009d\u009e"+
		"\7t\2\2\u009e\u009f\7k\2\2\u009f\u00a0\7p\2\2\u00a0\u00a1\7i\2\2\u00a1"+
		"\f\3\2\2\2\u00a2\u00a3\7k\2\2\u00a3\u00a4\78\2\2\u00a4\u00a5\7\66\2\2"+
		"\u00a5\16\3\2\2\2\u00a6\u00a7\7h\2\2\u00a7\u00a8\78\2\2\u00a8\u00a9\7"+
		"\66\2\2\u00a9\20\3\2\2\2\u00aa\u00ab\7d\2\2\u00ab\u00ac\7q\2\2\u00ac\u00ad"+
		"\7q\2\2\u00ad\u00ae\7n\2\2\u00ae\22\3\2\2\2\u00af\u00b0\7e\2\2\u00b0\u00b1"+
		"\7j\2\2\u00b1\u00b2\7c\2\2\u00b2\u00b3\7t\2\2\u00b3\24\3\2\2\2\u00b4\u00b5"+
		"\7(\2\2\u00b5\u00b6\7u\2\2\u00b6\u00b7\7v\2\2\u00b7\u00b8\7t\2\2\u00b8"+
		"\26\3\2\2\2\u00b9\u00ba\7c\2\2\u00ba\u00bb\7u\2\2\u00bb\30\3\2\2\2\u00bc"+
		"\u00bd\7k\2\2\u00bd\u00be\7h\2\2\u00be\32\3\2\2\2\u00bf\u00c0\7g\2\2\u00c0"+
		"\u00c1\7n\2\2\u00c1\u00c2\7u\2\2\u00c2\u00c3\7g\2\2\u00c3\34\3\2\2\2\u00c4"+
		"\u00c5\7o\2\2\u00c5\u00c6\7c\2\2\u00c6\u00c7\7v\2\2\u00c7\u00c8\7e\2\2"+
		"\u00c8\u00c9\7j\2\2\u00c9\36\3\2\2\2\u00ca\u00cb\7y\2\2\u00cb\u00cc\7"+
		"j\2\2\u00cc\u00cd\7k\2\2\u00cd\u00ce\7n\2\2\u00ce\u00cf\7g\2\2\u00cf "+
		"\3\2\2\2\u00d0\u00d1\7d\2\2\u00d1\u00d2\7t\2\2\u00d2\u00d3\7g\2\2\u00d3"+
		"\u00d4\7c\2\2\u00d4\u00d5\7m\2\2\u00d5\"\3\2\2\2\u00d6\u00d7\7e\2\2\u00d7"+
		"\u00d8\7q\2\2\u00d8\u00d9\7p\2\2\u00d9\u00da\7v\2\2\u00da\u00db\7k\2\2"+
		"\u00db\u00dc\7p\2\2\u00dc\u00dd\7w\2\2\u00dd\u00de\7g\2\2\u00de$\3\2\2"+
		"\2\u00df\u00e0\7t\2\2\u00e0\u00e1\7g\2\2\u00e1\u00e2\7v\2\2\u00e2\u00e3"+
		"\7w\2\2\u00e3\u00e4\7t\2\2\u00e4\u00e5\7p\2\2\u00e5&\3\2\2\2\u00e6\u00e7"+
		"\7r\2\2\u00e7\u00e8\7q\2\2\u00e8\u00e9\7y\2\2\u00e9(\3\2\2\2\u00ea\u00eb"+
		"\7r\2\2\u00eb\u00ec\7q\2\2\u00ec\u00ed\7y\2\2\u00ed\u00ee\7h\2\2\u00ee"+
		"*\3\2\2\2\u00ef\u00f0\7v\2\2\u00f0\u00f1\7q\2\2\u00f1\u00f2\7a\2\2\u00f2"+
		"\u00f3\7u\2\2\u00f3\u00f4\7v\2\2\u00f4\u00f5\7t\2\2\u00f5\u00f6\7k\2\2"+
		"\u00f6\u00f7\7p\2\2\u00f7\u00f8\7i\2\2\u00f8,\3\2\2\2\u00f9\u00fa\7v\2"+
		"\2\u00fa\u00fb\7q\2\2\u00fb\u00fc\7a\2\2\u00fc\u00fd\7q\2\2\u00fd\u00fe"+
		"\7y\2\2\u00fe\u00ff\7p\2\2\u00ff\u0100\7g\2\2\u0100\u0101\7f\2\2\u0101"+
		".\3\2\2\2\u0102\u0103\7o\2\2\u0103\u0104\7c\2\2\u0104\u0105\7k\2\2\u0105"+
		"\u0106\7p\2\2\u0106\60\3\2\2\2\u0107\u0108\7h\2\2\u0108\u0109\7p\2\2\u0109"+
		"\62\3\2\2\2\u010a\u010b\7\61\2\2\u010b\u010c\7,\2\2\u010c\u010e\3\2\2"+
		"\2\u010d\u010f\n\2\2\2\u010e\u010d\3\2\2\2\u010f\u0110\3\2\2\2\u0110\u010e"+
		"\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u0112\3\2\2\2\u0112\u0113\7,\2\2\u0113"+
		"\u0114\7\61\2\2\u0114\u0115\3\2\2\2\u0115\u0116\b\32\2\2\u0116\64\3\2"+
		"\2\2\u0117\u0118\7\61\2\2\u0118\u0119\7\61\2\2\u0119\u011b\3\2\2\2\u011a"+
		"\u011c\n\3\2\2\u011b\u011a\3\2\2\2\u011c\u011d\3\2\2\2\u011d\u011b\3\2"+
		"\2\2\u011d\u011e\3\2\2\2\u011e\u011f\3\2\2\2\u011f\u0120\b\33\2\2\u0120"+
		"\66\3\2\2\2\u0121\u0123\t\4\2\2\u0122\u0121\3\2\2\2\u0123\u0124\3\2\2"+
		"\2\u0124\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u01258\3\2\2\2\u0126\u0128"+
		"\t\4\2\2\u0127\u0126\3\2\2\2\u0128\u0129\3\2\2\2\u0129\u0127\3\2\2\2\u0129"+
		"\u012a\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u012d\7\60\2\2\u012c\u012e\t"+
		"\4\2\2\u012d\u012c\3\2\2\2\u012e\u012f\3\2\2\2\u012f\u012d\3\2\2\2\u012f"+
		"\u0130\3\2\2\2\u0130:\3\2\2\2\u0131\u0132\7)\2\2\u0132\u0133\n\5\2\2\u0133"+
		"\u0134\7)\2\2\u0134<\3\2\2\2\u0135\u0139\7$\2\2\u0136\u0138\n\5\2\2\u0137"+
		"\u0136\3\2\2\2\u0138\u013b\3\2\2\2\u0139\u0137\3\2\2\2\u0139\u013a\3\2"+
		"\2\2\u013a\u013c\3\2\2\2\u013b\u0139\3\2\2\2\u013c\u013d\7$\2\2\u013d"+
		">\3\2\2\2\u013e\u013f\7v\2\2\u013f\u0140\7t\2\2\u0140\u0141\7w\2\2\u0141"+
		"\u0148\7g\2\2\u0142\u0143\7h\2\2\u0143\u0144\7c\2\2\u0144\u0145\7n\2\2"+
		"\u0145\u0146\7u\2\2\u0146\u0148\7g\2\2\u0147\u013e\3\2\2\2\u0147\u0142"+
		"\3\2\2\2\u0148@\3\2\2\2\u0149\u014d\t\6\2\2\u014a\u014c\t\7\2\2\u014b"+
		"\u014a\3\2\2\2\u014c\u014f\3\2\2\2\u014d\u014b\3\2\2\2\u014d\u014e\3\2"+
		"\2\2\u014eB\3\2\2\2\u014f\u014d\3\2\2\2\u0150\u0151\7?\2\2\u0151D\3\2"+
		"\2\2\u0152\u0153\7\60\2\2\u0153F\3\2\2\2\u0154\u0155\7=\2\2\u0155H\3\2"+
		"\2\2\u0156\u0157\7.\2\2\u0157J\3\2\2\2\u0158\u0159\7<\2\2\u0159L\3\2\2"+
		"\2\u015a\u015b\7#\2\2\u015bN\3\2\2\2\u015c\u015d\7(\2\2\u015dP\3\2\2\2"+
		"\u015e\u015f\7<\2\2\u015f\u0160\7<\2\2\u0160R\3\2\2\2\u0161\u0162\7/\2"+
		"\2\u0162\u0163\7@\2\2\u0163T\3\2\2\2\u0164\u0165\7*\2\2\u0165V\3\2\2\2"+
		"\u0166\u0167\7+\2\2\u0167X\3\2\2\2\u0168\u0169\7}\2\2\u0169Z\3\2\2\2\u016a"+
		"\u016b\7\177\2\2\u016b\\\3\2\2\2\u016c\u016d\7]\2\2\u016d^\3\2\2\2\u016e"+
		"\u016f\7_\2\2\u016f`\3\2\2\2\u0170\u0171\7?\2\2\u0171\u0172\7?\2\2\u0172"+
		"b\3\2\2\2\u0173\u0174\7#\2\2\u0174\u0175\7?\2\2\u0175d\3\2\2\2\u0176\u0177"+
		"\7@\2\2\u0177f\3\2\2\2\u0178\u0179\7>\2\2\u0179h\3\2\2\2\u017a\u017b\7"+
		"@\2\2\u017b\u017c\7?\2\2\u017cj\3\2\2\2\u017d\u017e\7>\2\2\u017e\u017f"+
		"\7?\2\2\u017fl\3\2\2\2\u0180\u0181\7,\2\2\u0181n\3\2\2\2\u0182\u0183\7"+
		"\61\2\2\u0183p\3\2\2\2\u0184\u0185\7\'\2\2\u0185r\3\2\2\2\u0186\u0187"+
		"\7-\2\2\u0187t\3\2\2\2\u0188\u0189\7/\2\2\u0189v\3\2\2\2\u018a\u018b\7"+
		"(\2\2\u018b\u018c\7(\2\2\u018cx\3\2\2\2\u018d\u018e\7~\2\2\u018e\u018f"+
		"\7~\2\2\u018fz\3\2\2\2\u0190\u0191\7~\2\2\u0191|\3\2\2\2\u0192\u0193\7"+
		"?\2\2\u0193\u0194\7@\2\2\u0194~\3\2\2\2\u0195\u0196\7a\2\2\u0196\u0080"+
		"\3\2\2\2\u0197\u0199\t\b\2\2\u0198\u0197\3\2\2\2\u0199\u019a\3\2\2\2\u019a"+
		"\u0198\3\2\2\2\u019a\u019b\3\2\2\2\u019b\u019c\3\2\2\2\u019c\u019d\bA"+
		"\2\2\u019d\u0082\3\2\2\2\u019e\u019f\7^\2\2\u019f\u01a0\t\t\2\2\u01a0"+
		"\u0084\3\2\2\2\f\2\u0110\u011d\u0124\u0129\u012f\u0139\u0147\u014d\u019a"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}