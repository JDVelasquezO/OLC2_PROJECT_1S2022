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
		RLOOP=16, RBREAK=17, RCONTINUE=18, RRETURN=19, POWI=20, POWF=21, TOSTRING=22, 
		TOOWNED=23, RMAIN=24, RFN=25, MULTILINE=26, INLINE=27, INTEGER=28, FLOAT=29, 
		CHAR=30, STRING=31, BOOLEAN=32, ID=33, EQUAL=34, DOT=35, SEMICOLON=36, 
		COMMA=37, COLON=38, ADMIRATION=39, REFERENCE=40, HERITAGE=41, ARROW=42, 
		LEFT_PAR=43, RIGHT_PAR=44, LEFT_KEY=45, RIGHT_KEY=46, LEFT_BRACKET=47, 
		RIGHT_BRACKET=48, EQUEAL_EQUAL=49, NOTEQUAL=50, GREATER_THAN=51, LESS_THAN=52, 
		GREATER_EQUALTHAN=53, LESS_EQUEALTHAN=54, MUL=55, DIV=56, MOD=57, ADD=58, 
		SUB=59, AND=60, OR=61, PIPE=62, EQUAL_ARROW=63, UNDERSCORE=64, WHITESPACE=65;
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
			"RLOOP", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", "TOSTRING", 
			"TOOWNED", "RMAIN", "RFN", "MULTILINE", "INLINE", "INTEGER", "FLOAT", 
			"CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", 
			"COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "ARROW", "LEFT_PAR", 
			"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", 
			"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", 
			"LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "PIPE", 
			"EQUAL_ARROW", "UNDERSCORE", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", 
			"'f64'", "'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'match'", 
			"'while'", "'loop'", "'break'", "'continue'", "'return'", "'pow'", "'powf'", 
			"'to_string'", "'to_owned'", "'main'", "'fn'", null, null, null, null, 
			null, null, null, null, "'='", "'.'", "';'", "','", "':'", "'!'", "'&'", 
			"'::'", "'->'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", 
			"'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'", "'&&'", 
			"'||'", "'|'", "'=>'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE", 
			"RLOOP", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF", "TOSTRING", 
			"TOOWNED", "RMAIN", "RFN", "MULTILINE", "INLINE", "INTEGER", "FLOAT", 
			"CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", 
			"COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "ARROW", "LEFT_PAR", 
			"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", 
			"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", 
			"LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "PIPE", 
			"EQUAL_ARROW", "UNDERSCORE", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2C\u01a8\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3"+
		"\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3"+
		"\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3"+
		"\20\3\20\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3"+
		"\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3"+
		"\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\33\3\33\3\33\3\33\6\33\u0116"+
		"\n\33\r\33\16\33\u0117\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\6"+
		"\34\u0123\n\34\r\34\16\34\u0124\3\34\3\34\3\35\6\35\u012a\n\35\r\35\16"+
		"\35\u012b\3\36\6\36\u012f\n\36\r\36\16\36\u0130\3\36\3\36\6\36\u0135\n"+
		"\36\r\36\16\36\u0136\3\37\3\37\3\37\3\37\3 \3 \7 \u013f\n \f \16 \u0142"+
		"\13 \3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3!\5!\u014f\n!\3\"\3\"\7\"\u0153\n"+
		"\"\f\"\16\"\u0156\13\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*"+
		"\3*\3*\3+\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62"+
		"\3\62\3\63\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\66\3\67\3\67\3\67"+
		"\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3=\3>\3>\3>\3?\3?\3@\3@\3@\3A\3A"+
		"\3B\6B\u01a0\nB\rB\16B\u01a1\3B\3B\3C\3C\3C\2\2D\3\3\5\4\7\5\t\6\13\7"+
		"\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081"+
		"B\u0083C\u0085\2\3\2\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7\2C\\aac|\u00d3"+
		"\u00d3\u00f3\u00f3\b\2\62;C\\aac|\u00d3\u00d3\u00f3\u00f3\5\2\13\f\17"+
		"\17\"\"\t\2\"#%%--/\60<<BB]_\2\u01af\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2"+
		"\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23"+
		"\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2"+
		"\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2"+
		"\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3"+
		"\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2"+
		"\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2"+
		"\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2["+
		"\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2"+
		"\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2"+
		"\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2"+
		"\2\u0081\3\2\2\2\2\u0083\3\2\2\2\3\u0087\3\2\2\2\5\u008f\3\2\2\2\7\u0095"+
		"\3\2\2\2\t\u0099\3\2\2\2\13\u009d\3\2\2\2\r\u00a4\3\2\2\2\17\u00a8\3\2"+
		"\2\2\21\u00ac\3\2\2\2\23\u00b1\3\2\2\2\25\u00b6\3\2\2\2\27\u00bb\3\2\2"+
		"\2\31\u00be\3\2\2\2\33\u00c1\3\2\2\2\35\u00c6\3\2\2\2\37\u00cc\3\2\2\2"+
		"!\u00d2\3\2\2\2#\u00d7\3\2\2\2%\u00dd\3\2\2\2\'\u00e6\3\2\2\2)\u00ed\3"+
		"\2\2\2+\u00f1\3\2\2\2-\u00f6\3\2\2\2/\u0100\3\2\2\2\61\u0109\3\2\2\2\63"+
		"\u010e\3\2\2\2\65\u0111\3\2\2\2\67\u011e\3\2\2\29\u0129\3\2\2\2;\u012e"+
		"\3\2\2\2=\u0138\3\2\2\2?\u013c\3\2\2\2A\u014e\3\2\2\2C\u0150\3\2\2\2E"+
		"\u0157\3\2\2\2G\u0159\3\2\2\2I\u015b\3\2\2\2K\u015d\3\2\2\2M\u015f\3\2"+
		"\2\2O\u0161\3\2\2\2Q\u0163\3\2\2\2S\u0165\3\2\2\2U\u0168\3\2\2\2W\u016b"+
		"\3\2\2\2Y\u016d\3\2\2\2[\u016f\3\2\2\2]\u0171\3\2\2\2_\u0173\3\2\2\2a"+
		"\u0175\3\2\2\2c\u0177\3\2\2\2e\u017a\3\2\2\2g\u017d\3\2\2\2i\u017f\3\2"+
		"\2\2k\u0181\3\2\2\2m\u0184\3\2\2\2o\u0187\3\2\2\2q\u0189\3\2\2\2s\u018b"+
		"\3\2\2\2u\u018d\3\2\2\2w\u018f\3\2\2\2y\u0191\3\2\2\2{\u0194\3\2\2\2}"+
		"\u0197\3\2\2\2\177\u0199\3\2\2\2\u0081\u019c\3\2\2\2\u0083\u019f\3\2\2"+
		"\2\u0085\u01a5\3\2\2\2\u0087\u0088\7r\2\2\u0088\u0089\7t\2\2\u0089\u008a"+
		"\7k\2\2\u008a\u008b\7p\2\2\u008b\u008c\7v\2\2\u008c\u008d\7n\2\2\u008d"+
		"\u008e\7p\2\2\u008e\4\3\2\2\2\u008f\u0090\7r\2\2\u0090\u0091\7t\2\2\u0091"+
		"\u0092\7k\2\2\u0092\u0093\7p\2\2\u0093\u0094\7v\2\2\u0094\6\3\2\2\2\u0095"+
		"\u0096\7n\2\2\u0096\u0097\7g\2\2\u0097\u0098\7v\2\2\u0098\b\3\2\2\2\u0099"+
		"\u009a\7o\2\2\u009a\u009b\7w\2\2\u009b\u009c\7v\2\2\u009c\n\3\2\2\2\u009d"+
		"\u009e\7U\2\2\u009e\u009f\7v\2\2\u009f\u00a0\7t\2\2\u00a0\u00a1\7k\2\2"+
		"\u00a1\u00a2\7p\2\2\u00a2\u00a3\7i\2\2\u00a3\f\3\2\2\2\u00a4\u00a5\7k"+
		"\2\2\u00a5\u00a6\78\2\2\u00a6\u00a7\7\66\2\2\u00a7\16\3\2\2\2\u00a8\u00a9"+
		"\7h\2\2\u00a9\u00aa\78\2\2\u00aa\u00ab\7\66\2\2\u00ab\20\3\2\2\2\u00ac"+
		"\u00ad\7d\2\2\u00ad\u00ae\7q\2\2\u00ae\u00af\7q\2\2\u00af\u00b0\7n\2\2"+
		"\u00b0\22\3\2\2\2\u00b1\u00b2\7e\2\2\u00b2\u00b3\7j\2\2\u00b3\u00b4\7"+
		"c\2\2\u00b4\u00b5\7t\2\2\u00b5\24\3\2\2\2\u00b6\u00b7\7(\2\2\u00b7\u00b8"+
		"\7u\2\2\u00b8\u00b9\7v\2\2\u00b9\u00ba\7t\2\2\u00ba\26\3\2\2\2\u00bb\u00bc"+
		"\7c\2\2\u00bc\u00bd\7u\2\2\u00bd\30\3\2\2\2\u00be\u00bf\7k\2\2\u00bf\u00c0"+
		"\7h\2\2\u00c0\32\3\2\2\2\u00c1\u00c2\7g\2\2\u00c2\u00c3\7n\2\2\u00c3\u00c4"+
		"\7u\2\2\u00c4\u00c5\7g\2\2\u00c5\34\3\2\2\2\u00c6\u00c7\7o\2\2\u00c7\u00c8"+
		"\7c\2\2\u00c8\u00c9\7v\2\2\u00c9\u00ca\7e\2\2\u00ca\u00cb\7j\2\2\u00cb"+
		"\36\3\2\2\2\u00cc\u00cd\7y\2\2\u00cd\u00ce\7j\2\2\u00ce\u00cf\7k\2\2\u00cf"+
		"\u00d0\7n\2\2\u00d0\u00d1\7g\2\2\u00d1 \3\2\2\2\u00d2\u00d3\7n\2\2\u00d3"+
		"\u00d4\7q\2\2\u00d4\u00d5\7q\2\2\u00d5\u00d6\7r\2\2\u00d6\"\3\2\2\2\u00d7"+
		"\u00d8\7d\2\2\u00d8\u00d9\7t\2\2\u00d9\u00da\7g\2\2\u00da\u00db\7c\2\2"+
		"\u00db\u00dc\7m\2\2\u00dc$\3\2\2\2\u00dd\u00de\7e\2\2\u00de\u00df\7q\2"+
		"\2\u00df\u00e0\7p\2\2\u00e0\u00e1\7v\2\2\u00e1\u00e2\7k\2\2\u00e2\u00e3"+
		"\7p\2\2\u00e3\u00e4\7w\2\2\u00e4\u00e5\7g\2\2\u00e5&\3\2\2\2\u00e6\u00e7"+
		"\7t\2\2\u00e7\u00e8\7g\2\2\u00e8\u00e9\7v\2\2\u00e9\u00ea\7w\2\2\u00ea"+
		"\u00eb\7t\2\2\u00eb\u00ec\7p\2\2\u00ec(\3\2\2\2\u00ed\u00ee\7r\2\2\u00ee"+
		"\u00ef\7q\2\2\u00ef\u00f0\7y\2\2\u00f0*\3\2\2\2\u00f1\u00f2\7r\2\2\u00f2"+
		"\u00f3\7q\2\2\u00f3\u00f4\7y\2\2\u00f4\u00f5\7h\2\2\u00f5,\3\2\2\2\u00f6"+
		"\u00f7\7v\2\2\u00f7\u00f8\7q\2\2\u00f8\u00f9\7a\2\2\u00f9\u00fa\7u\2\2"+
		"\u00fa\u00fb\7v\2\2\u00fb\u00fc\7t\2\2\u00fc\u00fd\7k\2\2\u00fd\u00fe"+
		"\7p\2\2\u00fe\u00ff\7i\2\2\u00ff.\3\2\2\2\u0100\u0101\7v\2\2\u0101\u0102"+
		"\7q\2\2\u0102\u0103\7a\2\2\u0103\u0104\7q\2\2\u0104\u0105\7y\2\2\u0105"+
		"\u0106\7p\2\2\u0106\u0107\7g\2\2\u0107\u0108\7f\2\2\u0108\60\3\2\2\2\u0109"+
		"\u010a\7o\2\2\u010a\u010b\7c\2\2\u010b\u010c\7k\2\2\u010c\u010d\7p\2\2"+
		"\u010d\62\3\2\2\2\u010e\u010f\7h\2\2\u010f\u0110\7p\2\2\u0110\64\3\2\2"+
		"\2\u0111\u0112\7\61\2\2\u0112\u0113\7,\2\2\u0113\u0115\3\2\2\2\u0114\u0116"+
		"\n\2\2\2\u0115\u0114\3\2\2\2\u0116\u0117\3\2\2\2\u0117\u0115\3\2\2\2\u0117"+
		"\u0118\3\2\2\2\u0118\u0119\3\2\2\2\u0119\u011a\7,\2\2\u011a\u011b\7\61"+
		"\2\2\u011b\u011c\3\2\2\2\u011c\u011d\b\33\2\2\u011d\66\3\2\2\2\u011e\u011f"+
		"\7\61\2\2\u011f\u0120\7\61\2\2\u0120\u0122\3\2\2\2\u0121\u0123\n\3\2\2"+
		"\u0122\u0121\3\2\2\2\u0123\u0124\3\2\2\2\u0124\u0122\3\2\2\2\u0124\u0125"+
		"\3\2\2\2\u0125\u0126\3\2\2\2\u0126\u0127\b\34\2\2\u01278\3\2\2\2\u0128"+
		"\u012a\t\4\2\2\u0129\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u0129\3\2"+
		"\2\2\u012b\u012c\3\2\2\2\u012c:\3\2\2\2\u012d\u012f\t\4\2\2\u012e\u012d"+
		"\3\2\2\2\u012f\u0130\3\2\2\2\u0130\u012e\3\2\2\2\u0130\u0131\3\2\2\2\u0131"+
		"\u0132\3\2\2\2\u0132\u0134\7\60\2\2\u0133\u0135\t\4\2\2\u0134\u0133\3"+
		"\2\2\2\u0135\u0136\3\2\2\2\u0136\u0134\3\2\2\2\u0136\u0137\3\2\2\2\u0137"+
		"<\3\2\2\2\u0138\u0139\7)\2\2\u0139\u013a\n\5\2\2\u013a\u013b\7)\2\2\u013b"+
		">\3\2\2\2\u013c\u0140\7$\2\2\u013d\u013f\n\5\2\2\u013e\u013d\3\2\2\2\u013f"+
		"\u0142\3\2\2\2\u0140\u013e\3\2\2\2\u0140\u0141\3\2\2\2\u0141\u0143\3\2"+
		"\2\2\u0142\u0140\3\2\2\2\u0143\u0144\7$\2\2\u0144@\3\2\2\2\u0145\u0146"+
		"\7v\2\2\u0146\u0147\7t\2\2\u0147\u0148\7w\2\2\u0148\u014f\7g\2\2\u0149"+
		"\u014a\7h\2\2\u014a\u014b\7c\2\2\u014b\u014c\7n\2\2\u014c\u014d\7u\2\2"+
		"\u014d\u014f\7g\2\2\u014e\u0145\3\2\2\2\u014e\u0149\3\2\2\2\u014fB\3\2"+
		"\2\2\u0150\u0154\t\6\2\2\u0151\u0153\t\7\2\2\u0152\u0151\3\2\2\2\u0153"+
		"\u0156\3\2\2\2\u0154\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155D\3\2\2\2"+
		"\u0156\u0154\3\2\2\2\u0157\u0158\7?\2\2\u0158F\3\2\2\2\u0159\u015a\7\60"+
		"\2\2\u015aH\3\2\2\2\u015b\u015c\7=\2\2\u015cJ\3\2\2\2\u015d\u015e\7.\2"+
		"\2\u015eL\3\2\2\2\u015f\u0160\7<\2\2\u0160N\3\2\2\2\u0161\u0162\7#\2\2"+
		"\u0162P\3\2\2\2\u0163\u0164\7(\2\2\u0164R\3\2\2\2\u0165\u0166\7<\2\2\u0166"+
		"\u0167\7<\2\2\u0167T\3\2\2\2\u0168\u0169\7/\2\2\u0169\u016a\7@\2\2\u016a"+
		"V\3\2\2\2\u016b\u016c\7*\2\2\u016cX\3\2\2\2\u016d\u016e\7+\2\2\u016eZ"+
		"\3\2\2\2\u016f\u0170\7}\2\2\u0170\\\3\2\2\2\u0171\u0172\7\177\2\2\u0172"+
		"^\3\2\2\2\u0173\u0174\7]\2\2\u0174`\3\2\2\2\u0175\u0176\7_\2\2\u0176b"+
		"\3\2\2\2\u0177\u0178\7?\2\2\u0178\u0179\7?\2\2\u0179d\3\2\2\2\u017a\u017b"+
		"\7#\2\2\u017b\u017c\7?\2\2\u017cf\3\2\2\2\u017d\u017e\7@\2\2\u017eh\3"+
		"\2\2\2\u017f\u0180\7>\2\2\u0180j\3\2\2\2\u0181\u0182\7@\2\2\u0182\u0183"+
		"\7?\2\2\u0183l\3\2\2\2\u0184\u0185\7>\2\2\u0185\u0186\7?\2\2\u0186n\3"+
		"\2\2\2\u0187\u0188\7,\2\2\u0188p\3\2\2\2\u0189\u018a\7\61\2\2\u018ar\3"+
		"\2\2\2\u018b\u018c\7\'\2\2\u018ct\3\2\2\2\u018d\u018e\7-\2\2\u018ev\3"+
		"\2\2\2\u018f\u0190\7/\2\2\u0190x\3\2\2\2\u0191\u0192\7(\2\2\u0192\u0193"+
		"\7(\2\2\u0193z\3\2\2\2\u0194\u0195\7~\2\2\u0195\u0196\7~\2\2\u0196|\3"+
		"\2\2\2\u0197\u0198\7~\2\2\u0198~\3\2\2\2\u0199\u019a\7?\2\2\u019a\u019b"+
		"\7@\2\2\u019b\u0080\3\2\2\2\u019c\u019d\7a\2\2\u019d\u0082\3\2\2\2\u019e"+
		"\u01a0\t\b\2\2\u019f\u019e\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1\u019f\3\2"+
		"\2\2\u01a1\u01a2\3\2\2\2\u01a2\u01a3\3\2\2\2\u01a3\u01a4\bB\2\2\u01a4"+
		"\u0084\3\2\2\2\u01a5\u01a6\7^\2\2\u01a6\u01a7\t\t\2\2\u01a7\u0086\3\2"+
		"\2\2\f\2\u0117\u0124\u012b\u0130\u0136\u0140\u014e\u0154\u01a1\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}