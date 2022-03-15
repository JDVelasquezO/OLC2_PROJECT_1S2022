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
		POWI=16, POWF=17, TOSTRING=18, TOOWNED=19, RMAIN=20, RFN=21, MULTILINE=22, 
		INLINE=23, INTEGER=24, FLOAT=25, CHAR=26, STRING=27, BOOLEAN=28, ID=29, 
		EQUAL=30, DOT=31, SEMICOLON=32, COMMA=33, COLON=34, ADMIRATION=35, REFERENCE=36, 
		HERITAGE=37, ARROW=38, LEFT_PAR=39, RIGHT_PAR=40, LEFT_KEY=41, RIGHT_KEY=42, 
		LEFT_BRACKET=43, RIGHT_BRACKET=44, EQUEAL_EQUAL=45, NOTEQUAL=46, GREATER_THAN=47, 
		LESS_THAN=48, GREATER_EQUALTHAN=49, LESS_EQUEALTHAN=50, MUL=51, DIV=52, 
		MOD=53, ADD=54, SUB=55, AND=56, OR=57, PIPE=58, EQUAL_ARROW=59, UNDERSCORE=60, 
		WHITESPACE=61;
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
			"POWI", "POWF", "TOSTRING", "TOOWNED", "RMAIN", "RFN", "MULTILINE", "INLINE", 
			"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", 
			"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", 
			"ARROW", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", 
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
			"'while'", "'pow'", "'powf'", "'to_string'", "'to_owned'", "'main'", 
			"'fn'", null, null, null, null, null, null, null, null, "'='", "'.'", 
			"';'", "','", "':'", "'!'", "'&'", "'::'", "'->'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'*'", 
			"'/'", "'%'", "'+'", "'-'", "'&&'", "'||'", "'|'", "'=>'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"POWI", "POWF", "TOSTRING", "TOOWNED", "RMAIN", "RFN", "MULTILINE", "INLINE", 
			"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", 
			"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", 
			"ARROW", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2?\u0185\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13"+
		"\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25"+
		"\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27\3\27\6\27\u00f3\n\27\r\27\16"+
		"\27\u00f4\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\6\30\u0100\n\30"+
		"\r\30\16\30\u0101\3\30\3\30\3\31\6\31\u0107\n\31\r\31\16\31\u0108\3\32"+
		"\6\32\u010c\n\32\r\32\16\32\u010d\3\32\3\32\6\32\u0112\n\32\r\32\16\32"+
		"\u0113\3\33\3\33\3\33\3\33\3\34\3\34\7\34\u011c\n\34\f\34\16\34\u011f"+
		"\13\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\5\35\u012c"+
		"\n\35\3\36\3\36\7\36\u0130\n\36\f\36\16\36\u0133\13\36\3\37\3\37\3 \3"+
		" \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3&\3\'\3\'\3\'\3(\3(\3)\3)\3*"+
		"\3*\3+\3+\3,\3,\3-\3-\3.\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62"+
		"\3\62\3\63\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39"+
		"\39\39\3:\3:\3:\3;\3;\3<\3<\3<\3=\3=\3>\6>\u017d\n>\r>\16>\u017e\3>\3"+
		">\3?\3?\3?\2\2@\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64"+
		"g\65i\66k\67m8o9q:s;u<w=y>{?}\2\3\2\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$"+
		"\7\2C\\aac|\u00d3\u00d3\u00f3\u00f3\b\2\62;C\\aac|\u00d3\u00d3\u00f3\u00f3"+
		"\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u018c\2\3\3\2\2\2\2\5\3\2\2"+
		"\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2"+
		"\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2"+
		"Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3"+
		"\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2"+
		"\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\3\177\3\2\2"+
		"\2\5\u0087\3\2\2\2\7\u008d\3\2\2\2\t\u0091\3\2\2\2\13\u0095\3\2\2\2\r"+
		"\u009c\3\2\2\2\17\u00a0\3\2\2\2\21\u00a4\3\2\2\2\23\u00a9\3\2\2\2\25\u00ae"+
		"\3\2\2\2\27\u00b3\3\2\2\2\31\u00b6\3\2\2\2\33\u00b9\3\2\2\2\35\u00be\3"+
		"\2\2\2\37\u00c4\3\2\2\2!\u00ca\3\2\2\2#\u00ce\3\2\2\2%\u00d3\3\2\2\2\'"+
		"\u00dd\3\2\2\2)\u00e6\3\2\2\2+\u00eb\3\2\2\2-\u00ee\3\2\2\2/\u00fb\3\2"+
		"\2\2\61\u0106\3\2\2\2\63\u010b\3\2\2\2\65\u0115\3\2\2\2\67\u0119\3\2\2"+
		"\29\u012b\3\2\2\2;\u012d\3\2\2\2=\u0134\3\2\2\2?\u0136\3\2\2\2A\u0138"+
		"\3\2\2\2C\u013a\3\2\2\2E\u013c\3\2\2\2G\u013e\3\2\2\2I\u0140\3\2\2\2K"+
		"\u0142\3\2\2\2M\u0145\3\2\2\2O\u0148\3\2\2\2Q\u014a\3\2\2\2S\u014c\3\2"+
		"\2\2U\u014e\3\2\2\2W\u0150\3\2\2\2Y\u0152\3\2\2\2[\u0154\3\2\2\2]\u0157"+
		"\3\2\2\2_\u015a\3\2\2\2a\u015c\3\2\2\2c\u015e\3\2\2\2e\u0161\3\2\2\2g"+
		"\u0164\3\2\2\2i\u0166\3\2\2\2k\u0168\3\2\2\2m\u016a\3\2\2\2o\u016c\3\2"+
		"\2\2q\u016e\3\2\2\2s\u0171\3\2\2\2u\u0174\3\2\2\2w\u0176\3\2\2\2y\u0179"+
		"\3\2\2\2{\u017c\3\2\2\2}\u0182\3\2\2\2\177\u0080\7r\2\2\u0080\u0081\7"+
		"t\2\2\u0081\u0082\7k\2\2\u0082\u0083\7p\2\2\u0083\u0084\7v\2\2\u0084\u0085"+
		"\7n\2\2\u0085\u0086\7p\2\2\u0086\4\3\2\2\2\u0087\u0088\7r\2\2\u0088\u0089"+
		"\7t\2\2\u0089\u008a\7k\2\2\u008a\u008b\7p\2\2\u008b\u008c\7v\2\2\u008c"+
		"\6\3\2\2\2\u008d\u008e\7n\2\2\u008e\u008f\7g\2\2\u008f\u0090\7v\2\2\u0090"+
		"\b\3\2\2\2\u0091\u0092\7o\2\2\u0092\u0093\7w\2\2\u0093\u0094\7v\2\2\u0094"+
		"\n\3\2\2\2\u0095\u0096\7U\2\2\u0096\u0097\7v\2\2\u0097\u0098\7t\2\2\u0098"+
		"\u0099\7k\2\2\u0099\u009a\7p\2\2\u009a\u009b\7i\2\2\u009b\f\3\2\2\2\u009c"+
		"\u009d\7k\2\2\u009d\u009e\78\2\2\u009e\u009f\7\66\2\2\u009f\16\3\2\2\2"+
		"\u00a0\u00a1\7h\2\2\u00a1\u00a2\78\2\2\u00a2\u00a3\7\66\2\2\u00a3\20\3"+
		"\2\2\2\u00a4\u00a5\7d\2\2\u00a5\u00a6\7q\2\2\u00a6\u00a7\7q\2\2\u00a7"+
		"\u00a8\7n\2\2\u00a8\22\3\2\2\2\u00a9\u00aa\7e\2\2\u00aa\u00ab\7j\2\2\u00ab"+
		"\u00ac\7c\2\2\u00ac\u00ad\7t\2\2\u00ad\24\3\2\2\2\u00ae\u00af\7(\2\2\u00af"+
		"\u00b0\7u\2\2\u00b0\u00b1\7v\2\2\u00b1\u00b2\7t\2\2\u00b2\26\3\2\2\2\u00b3"+
		"\u00b4\7c\2\2\u00b4\u00b5\7u\2\2\u00b5\30\3\2\2\2\u00b6\u00b7\7k\2\2\u00b7"+
		"\u00b8\7h\2\2\u00b8\32\3\2\2\2\u00b9\u00ba\7g\2\2\u00ba\u00bb\7n\2\2\u00bb"+
		"\u00bc\7u\2\2\u00bc\u00bd\7g\2\2\u00bd\34\3\2\2\2\u00be\u00bf\7o\2\2\u00bf"+
		"\u00c0\7c\2\2\u00c0\u00c1\7v\2\2\u00c1\u00c2\7e\2\2\u00c2\u00c3\7j\2\2"+
		"\u00c3\36\3\2\2\2\u00c4\u00c5\7y\2\2\u00c5\u00c6\7j\2\2\u00c6\u00c7\7"+
		"k\2\2\u00c7\u00c8\7n\2\2\u00c8\u00c9\7g\2\2\u00c9 \3\2\2\2\u00ca\u00cb"+
		"\7r\2\2\u00cb\u00cc\7q\2\2\u00cc\u00cd\7y\2\2\u00cd\"\3\2\2\2\u00ce\u00cf"+
		"\7r\2\2\u00cf\u00d0\7q\2\2\u00d0\u00d1\7y\2\2\u00d1\u00d2\7h\2\2\u00d2"+
		"$\3\2\2\2\u00d3\u00d4\7v\2\2\u00d4\u00d5\7q\2\2\u00d5\u00d6\7a\2\2\u00d6"+
		"\u00d7\7u\2\2\u00d7\u00d8\7v\2\2\u00d8\u00d9\7t\2\2\u00d9\u00da\7k\2\2"+
		"\u00da\u00db\7p\2\2\u00db\u00dc\7i\2\2\u00dc&\3\2\2\2\u00dd\u00de\7v\2"+
		"\2\u00de\u00df\7q\2\2\u00df\u00e0\7a\2\2\u00e0\u00e1\7q\2\2\u00e1\u00e2"+
		"\7y\2\2\u00e2\u00e3\7p\2\2\u00e3\u00e4\7g\2\2\u00e4\u00e5\7f\2\2\u00e5"+
		"(\3\2\2\2\u00e6\u00e7\7o\2\2\u00e7\u00e8\7c\2\2\u00e8\u00e9\7k\2\2\u00e9"+
		"\u00ea\7p\2\2\u00ea*\3\2\2\2\u00eb\u00ec\7h\2\2\u00ec\u00ed\7p\2\2\u00ed"+
		",\3\2\2\2\u00ee\u00ef\7\61\2\2\u00ef\u00f0\7,\2\2\u00f0\u00f2\3\2\2\2"+
		"\u00f1\u00f3\n\2\2\2\u00f2\u00f1\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f2"+
		"\3\2\2\2\u00f4\u00f5\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6\u00f7\7,\2\2\u00f7"+
		"\u00f8\7\61\2\2\u00f8\u00f9\3\2\2\2\u00f9\u00fa\b\27\2\2\u00fa.\3\2\2"+
		"\2\u00fb\u00fc\7\61\2\2\u00fc\u00fd\7\61\2\2\u00fd\u00ff\3\2\2\2\u00fe"+
		"\u0100\n\3\2\2\u00ff\u00fe\3\2\2\2\u0100\u0101\3\2\2\2\u0101\u00ff\3\2"+
		"\2\2\u0101\u0102\3\2\2\2\u0102\u0103\3\2\2\2\u0103\u0104\b\30\2\2\u0104"+
		"\60\3\2\2\2\u0105\u0107\t\4\2\2\u0106\u0105\3\2\2\2\u0107\u0108\3\2\2"+
		"\2\u0108\u0106\3\2\2\2\u0108\u0109\3\2\2\2\u0109\62\3\2\2\2\u010a\u010c"+
		"\t\4\2\2\u010b\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u010b\3\2\2\2\u010d"+
		"\u010e\3\2\2\2\u010e\u010f\3\2\2\2\u010f\u0111\7\60\2\2\u0110\u0112\t"+
		"\4\2\2\u0111\u0110\3\2\2\2\u0112\u0113\3\2\2\2\u0113\u0111\3\2\2\2\u0113"+
		"\u0114\3\2\2\2\u0114\64\3\2\2\2\u0115\u0116\7)\2\2\u0116\u0117\n\5\2\2"+
		"\u0117\u0118\7)\2\2\u0118\66\3\2\2\2\u0119\u011d\7$\2\2\u011a\u011c\n"+
		"\5\2\2\u011b\u011a\3\2\2\2\u011c\u011f\3\2\2\2\u011d\u011b\3\2\2\2\u011d"+
		"\u011e\3\2\2\2\u011e\u0120\3\2\2\2\u011f\u011d\3\2\2\2\u0120\u0121\7$"+
		"\2\2\u01218\3\2\2\2\u0122\u0123\7v\2\2\u0123\u0124\7t\2\2\u0124\u0125"+
		"\7w\2\2\u0125\u012c\7g\2\2\u0126\u0127\7h\2\2\u0127\u0128\7c\2\2\u0128"+
		"\u0129\7n\2\2\u0129\u012a\7u\2\2\u012a\u012c\7g\2\2\u012b\u0122\3\2\2"+
		"\2\u012b\u0126\3\2\2\2\u012c:\3\2\2\2\u012d\u0131\t\6\2\2\u012e\u0130"+
		"\t\7\2\2\u012f\u012e\3\2\2\2\u0130\u0133\3\2\2\2\u0131\u012f\3\2\2\2\u0131"+
		"\u0132\3\2\2\2\u0132<\3\2\2\2\u0133\u0131\3\2\2\2\u0134\u0135\7?\2\2\u0135"+
		">\3\2\2\2\u0136\u0137\7\60\2\2\u0137@\3\2\2\2\u0138\u0139\7=\2\2\u0139"+
		"B\3\2\2\2\u013a\u013b\7.\2\2\u013bD\3\2\2\2\u013c\u013d\7<\2\2\u013dF"+
		"\3\2\2\2\u013e\u013f\7#\2\2\u013fH\3\2\2\2\u0140\u0141\7(\2\2\u0141J\3"+
		"\2\2\2\u0142\u0143\7<\2\2\u0143\u0144\7<\2\2\u0144L\3\2\2\2\u0145\u0146"+
		"\7/\2\2\u0146\u0147\7@\2\2\u0147N\3\2\2\2\u0148\u0149\7*\2\2\u0149P\3"+
		"\2\2\2\u014a\u014b\7+\2\2\u014bR\3\2\2\2\u014c\u014d\7}\2\2\u014dT\3\2"+
		"\2\2\u014e\u014f\7\177\2\2\u014fV\3\2\2\2\u0150\u0151\7]\2\2\u0151X\3"+
		"\2\2\2\u0152\u0153\7_\2\2\u0153Z\3\2\2\2\u0154\u0155\7?\2\2\u0155\u0156"+
		"\7?\2\2\u0156\\\3\2\2\2\u0157\u0158\7#\2\2\u0158\u0159\7?\2\2\u0159^\3"+
		"\2\2\2\u015a\u015b\7@\2\2\u015b`\3\2\2\2\u015c\u015d\7>\2\2\u015db\3\2"+
		"\2\2\u015e\u015f\7@\2\2\u015f\u0160\7?\2\2\u0160d\3\2\2\2\u0161\u0162"+
		"\7>\2\2\u0162\u0163\7?\2\2\u0163f\3\2\2\2\u0164\u0165\7,\2\2\u0165h\3"+
		"\2\2\2\u0166\u0167\7\61\2\2\u0167j\3\2\2\2\u0168\u0169\7\'\2\2\u0169l"+
		"\3\2\2\2\u016a\u016b\7-\2\2\u016bn\3\2\2\2\u016c\u016d\7/\2\2\u016dp\3"+
		"\2\2\2\u016e\u016f\7(\2\2\u016f\u0170\7(\2\2\u0170r\3\2\2\2\u0171\u0172"+
		"\7~\2\2\u0172\u0173\7~\2\2\u0173t\3\2\2\2\u0174\u0175\7~\2\2\u0175v\3"+
		"\2\2\2\u0176\u0177\7?\2\2\u0177\u0178\7@\2\2\u0178x\3\2\2\2\u0179\u017a"+
		"\7a\2\2\u017az\3\2\2\2\u017b\u017d\t\b\2\2\u017c\u017b\3\2\2\2\u017d\u017e"+
		"\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017f\u0180\3\2\2\2\u0180"+
		"\u0181\b>\2\2\u0181|\3\2\2\2\u0182\u0183\7^\2\2\u0183\u0184\t\t\2\2\u0184"+
		"~\3\2\2\2\f\2\u00f4\u0101\u0108\u010d\u0113\u011d\u012b\u0131\u017e\3"+
		"\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}